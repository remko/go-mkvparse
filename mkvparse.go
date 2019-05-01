//go:generate go run generate.go

// Provides push-style parser functions for parsing Matroska (`.mkv`, `.mka`) files.
package mkvparse

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////

type ElementID int64

func (e ElementID) String() string {
	return fmt.Sprintf("%x", int64(e))
}

type ElementType int

const (
	_ = iota
	uintegerType
	integerType
	binaryType
	stringType
	utf8Type
	floatType
	dateType
	masterType
)

// https://github.com/cellar-wg/ebml-specification/blob/master/specification.markdown#element-data-size
//
// An Element Data Size with all VINT_DATA bits set to one is reserved as an
// indicator that the size of the EBML Element is unknown.
const UnknownSize = -1

type ElementInfo struct {
	Offset int64
	Size   int64
	Level  int
}

// Combined element id and information.
type Element struct {
	ID   ElementID
	Info ElementInfo
}

// Interface for handling parse events
type Handler interface {
	// Return `true` to descend into the element, `false` to skip this element's children.
	HandleMasterBegin(ElementID, ElementInfo) (bool, error)
	HandleMasterEnd(ElementID, ElementInfo) error

	HandleString(ElementID, string, ElementInfo) error
	HandleInteger(ElementID, int64, ElementInfo) error
	HandleFloat(ElementID, float64, ElementInfo) error
	HandleDate(ElementID, time.Time, ElementInfo) error
	HandleBinary(ElementID, []byte, ElementInfo) error

	HandleParseError(reader io.Reader, err *ParseError) *ParseError
}

// Parse the file pointed to by `path`.
func ParsePath(path string, handler Handler) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return Parse(file, handler)
}

// Parse the contents of `reader`.
func Parse(reader io.Reader, handler Handler) error {
	err := parse(reader, handler)

	if err != nil && err.Err != io.EOF {
		return err
	}

	return nil
}

func readN(r io.Reader, b []byte) (err error) {
	c, e := io.ReadAtLeast(r, b, len(b))
	if e == io.ErrUnexpectedEOF {
		err = &TruncatedInputError{
			Expected: int64(len(b)),
			Received: int64(c),
		}
		return err
	}

	return e
}

// Parse elements from the stream and call the handlers at the appropriate
// places. This will parse until there is an EOF or an error.
func parse(reader io.Reader, handler Handler) (perr *ParseError) {
	var offset int64

	stack := NewElementStack()

	// Helper for creation of parse errors and calling handler.
	eh := func(e error) *ParseError {
		if e != nil {
			pe := NewParseError(offset, stack, e)

			return handler.HandleParseError(reader, pe)
		}

		return nil
	}

	// Handle state when we reach an EOF condition.
	defer func() {
		if perr.Err == io.EOF || perr.Err == io.ErrUnexpectedEOF {
			for count := len(*stack); count > 0; count-- {
				tid, tinfo := stack.Peek()

				seen := offset - tinfo.Offset

				if tinfo.Size != UnknownSize && seen != tinfo.Size {
					err := &TruncatedInputError{
						Expected: tinfo.Size,
						Received: seen,
					}
					perr = eh(err)
					if perr != nil {
						return
					}
				}

				stack.Pop()

				err := handler.HandleMasterEnd(tid, tinfo)
				perr = eh(err)
				if perr != nil {
					return
				}
			}
		}
	}()

	// Continue parsing elements from the stream until it reaches EOF or
	// there is an error.
	for {
		id, idCount, err := readElementID(reader)
		perr = eh(err)
		if perr != nil {
			return perr
		}

		offset += idCount

		for len(*stack) > 0 {
			// Check that this is a valid child of the current
			// parent. If it isn't, then possibly this indicates an
			// UnknownSize parent has finished.
			pid, _ := stack.Peek()

			// A negative level is allowed as a child anywhere.
			if elementLevels[id] < 0 {
				break
			}

			_, ok := elementChildren[pid][id]
			if !ok {
				// Are any of the stack elements of unknown
				// size? If so, then unwind the stack to that
				// element. We don't break from the loop yet
				// because it is possible that multiple nested
				// UnknownSize master elements are ending at
				// the same time.
				i, ok := stack.UnknownSize()
				if ok {
					for count := len(*stack) - i; count > 0; count-- {
						tid, tinfo := stack.Pop()

						err := handler.HandleMasterEnd(tid, tinfo)
						perr = eh(err)
						if perr != nil {
							return perr
						}
					}
				} else {
					err := &InvalidElementError{
						Expected: elementChildren[pid],
						Received: id,
					}
					perr = eh(err)
					if perr != nil {
						return perr
					}
				}
			} else {
				break
			}
		}

		// Is this an invalid root element?
		if len(*stack) == 0 && elementLevels[id] > 0 {
			expected := map[ElementID]bool{}

			for k, v := range elementLevels {
				if v <= 0 {
					expected[k] = true
				}
			}

			err := &InvalidElementError{
				Expected: expected,
				Received: id,
			}
			perr = eh(err)
			if perr != nil {
				return perr
			}
		} else if len(*stack) > 0 {
			// Not a wildcard level?
			if elementLevels[id] >= 0 {
				// Is this a valid child of the current parent?
				pid, _ := stack.Peek()

				_, ok := elementChildren[pid][id]
				if !ok {
					err := &InvalidElementError{
						Expected: elementChildren[pid],
						Received: id,
					}
					perr = eh(err)
					if perr != nil {
						return perr
					}
				}
			}
		}

		size, sizeCount, err := readVarInt(reader)
		perr = eh(err)
		if perr != nil {
			return perr
		}

		offset += sizeCount

		info := ElementInfo{
			Offset: offset,
			Size:   size,
			Level:  len(*stack),
		}

		typ := elementTypes[id]

		if typ == masterType {
			stack.Push(id, info)

			descend, err := handler.HandleMasterBegin(id, info)
			perr = eh(err)
			if perr != nil {
				return perr
			}

			// Skip over the children if we don't want to descend.
			// This is currently only possible if the size is
			// known.
			if !descend && size != UnknownSize {
				data := make([]byte, size)

				err := readN(reader, data)
				perr = eh(err)
				if perr != nil {
					return perr
				}

				offset += size

				tid, tinfo := stack.Pop()

				err = handler.HandleMasterEnd(tid, tinfo)
				perr = eh(err)
				if perr != nil {
					return perr
				}
			}
		} else {
			// Temporarily push the leaf (non-master) element onto
			// the stack so that it is avaialble to the error
			// handler as the last element of the stack.
			stack.Push(id, info)

			data := make([]byte, size)

			err := readN(reader, data)
			perr = eh(err)
			if perr != nil {
				return perr
			}

			offset += size

			switch typ {
			case uintegerType:
				err = handler.HandleInteger(id, int64(binary.BigEndian.Uint64(pad(data, 8))), info)
			case integerType:
				err = handler.HandleInteger(id, convertBytesToSignedInt(data), info)
			case binaryType:
				err = handler.HandleBinary(id, data, info)
			case stringType, utf8Type:
				err = handler.HandleString(id, string(unpadString(data)), info)
			case floatType:
				var value float64

				switch size {
				case 0:
				case 4:
					value = float64(math.Float32frombits(binary.BigEndian.Uint32(data)))
				case 8:
					value = math.Float64frombits(binary.BigEndian.Uint64(data))
				default:
					perr = eh(&InvalidFloatSizeError{
						Expected: []int64{0, 4, 8},
						Received: size,
					})
					return perr
				}

				err = handler.HandleFloat(id, value, info)
			case dateType:
				err = handler.HandleDate(id, baseDate.Add(time.Duration(convertBytesToSignedInt(data))), info)
			}
			perr = eh(err)
			if perr != nil {
				return perr
			}

			stack.Pop()
		}

		// Check if we've seen enough data and pop if we have.
		for len(*stack) > 0 {
			_, pinfo := stack.Peek()

			seen := offset - pinfo.Offset

			if pinfo.Size != UnknownSize && seen >= pinfo.Size {
				tid, tinfo := stack.Pop()

				err = handler.HandleMasterEnd(tid, tinfo)
				perr = eh(err)
				if perr != nil {
					return perr
				}
			} else {
				break
			}
		}
	}

	return nil
}

// Gives the human-readable name for the given element ID.
func NameForElementID(id ElementID) string {
	name, ok := elementNames[id]
	if !ok {
		return "UNKNOWN"
	}
	return name
}

func ParentsForElementID(id ElementID) []ElementID {
	target := id
	parents := []ElementID{}

	for {
		parent, ok := elementParent[target]
		if !ok {
			break
		}

		parents = append(parents, parent)
		target = parent
	}

	return parents
}

////////////////////////////////////////////////////////////////////////////////
// Utility
////////////////////////////////////////////////////////////////////////////////

var baseDate = time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

func convertBytesToSignedInt(data []byte) int64 {
	if data[0] >= 0x80 {
		result := make([]byte, len(data))
		for i := range data {
			result[i] = ^data[i]
		}
		return -(int64(binary.BigEndian.Uint64(pad(result, 8))) + 1)
	} else {
		return int64(binary.BigEndian.Uint64(pad(data, 8)))
	}
}

func pad(b []byte, size int) []byte {
	if len(b) == size || len(b) > size {
		return b
	}
	tmp := make([]byte, size)
	copy(tmp[size-len(b):], b)
	return tmp
}

func unpadString(b []byte) []byte {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0x0 {
			return b[:i+1]
		}
	}
	return b[0:0]
}

func readVarInt(reader io.Reader) (int64, int64, error) {
	return readVarIntRaw(reader, true)
}

func readElementID(reader io.Reader) (ElementID, int64, error) {
	rawID, count, err := readVarIntRaw(reader, false)
	return ElementID(rawID), count, err
}

func readVarIntRaw(reader io.Reader, doMask bool) (int64, int64, error) {
	b := make([]byte, 1)

	err := readN(reader, b)
	if err != nil {
		return -1, -1, err
	}

	var length int
	var mask byte
	var unknown int64

	if ((b[0] & 0x80) >> 7) == 1 {
		length = 1
		mask = 0x7f
		unknown = 0x7f
	} else if ((b[0] & 0x40) >> 6) == 1 {
		length = 2
		mask = 0x3f
		unknown = 0x3fff
	} else if ((b[0] & 0x20) >> 5) == 1 {
		length = 3
		mask = 0x1f
		unknown = 0x1fffff
	} else if ((b[0] & 0x10) >> 4) == 1 {
		length = 4
		mask = 0xf
		unknown = 0x0fffffff
	} else if ((b[0] & 0x08) >> 3) == 1 {
		length = 5
		mask = 0x7
		unknown = 0x07ffffffff
	} else if ((b[0] & 0x04) >> 2) == 1 {
		length = 6
		mask = 0x3
		unknown = 0x03ffffffffff
	} else if ((b[0] & 0x02) >> 1) == 1 {
		length = 7
		mask = 0x1
		unknown = 0x01ffffffffffff
	} else if ((b[0] & 0x01) >> 0) == 1 {
		length = 8
		mask = 0x0
		unknown = 0x00ffffffffffffff
	}

	if length == 0 {
		return -1, -1, fmt.Errorf("Invalid length found: %x", b[0])
	}

	result := make([]byte, 8)

	if doMask {
		result[8-length] = b[0] & mask
	} else {
		result[8-length] = b[0]
	}

	err = readN(reader, result[8-length+1:])
	if err != nil {
		return -1, -1, err
	}

	val := int64(binary.BigEndian.Uint64(result))
	if val == unknown {
		val = UnknownSize
	}

	return val, int64(length), nil
}
