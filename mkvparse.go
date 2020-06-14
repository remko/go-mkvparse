//go:generate go run generate.go

// Package mkvparse provides push-style parser functions for parsing Matroska
// (`.mkv`, `.mka`) files.
package mkvparse

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////

// ElementID represents the EBML ID of an element.
// The supported EBML IDs are documented in the Matroska specification:
// https://www.matroska.org/technical/specs/index.html
type ElementID int64

type elementType int

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

// ElementInfo contains information about an element encountered in
// the stream, and is passed to the handler by the parser on parse events.
type ElementInfo struct {
	Offset int64
	Size   int64
	Level  int
}

// Handler declares an interface for handling parse events
type Handler interface {
	// Return `true` to descend into the element, `false` to skip this element's children.
	HandleMasterBegin(ElementID, ElementInfo) (bool, error)
	HandleMasterEnd(ElementID, ElementInfo) error
	HandleString(ElementID, string, ElementInfo) error
	HandleInteger(ElementID, int64, ElementInfo) error
	HandleFloat(ElementID, float64, ElementInfo) error
	HandleDate(ElementID, time.Time, ElementInfo) error
	HandleBinary(ElementID, []byte, ElementInfo) error
}

////////////////////////////////////////////////////////////////////////////////

// Parse the file pointed to by `path`
func ParsePath(path string, handler Handler) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return Parse(file, handler)
}

// Parse the contents of `reader`
func Parse(reader io.Reader, handler Handler) error {
	_, err := parseElements(reader, 0, -1, 0, handler)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// Parse all sibling elements on one level until 'size' bytes
// have been read (or until EOF)
func parseElements(reader io.Reader, currentOffset int64, size int64, level int, handler Handler) (count int64, err error) {
	for size < 0 || count < size {
		elementCount, err := parseElement(reader, currentOffset+count, level, handler)
		if err != nil {
			return -1, err
		}
		count = count + elementCount
	}
	return count, nil
}

// Parse one complete element.
// Recursively descends master elements.
func parseElement(reader io.Reader, currentOffset int64, level int, handler Handler) (count int64, err error) {
	id, idCount, err := readElementID(reader)
	if err != nil {
		return -1, err
	}
	size, sizeCount, err := readVarInt(reader)
	if err != nil {
		return -1, err
	}
	typ := elementTypes[id]
	// fmt.Printf("@%x %d %s %x %x\n", currentOffset, level, NameForElementID(id), size, typ)
	elementOffset := currentOffset + idCount + sizeCount
	count = idCount + sizeCount + size
	info := ElementInfo{
		Offset: elementOffset,
		Size:   size,
		Level:  level,
	}
	if typ == masterType {
		descend, err := handler.HandleMasterBegin(id, info)
		if err != nil {
			return -1, err
		}
		if descend {
			_, err := parseElements(reader, elementOffset, size, level+1, handler)
			if err != nil {
				return -1, err
			}
		} else {
			switch reader := reader.(type) {
			case io.Seeker:
				if _, err := reader.Seek(size, io.SeekCurrent); err != nil {
					return -1, err
				}
			default:
				if _, err := io.CopyN(ioutil.Discard, reader, size); err != nil {
					return -1, err
				}
			}
		}
		err = handler.HandleMasterEnd(id, info)
		if err != nil {
			return -1, err
		}
		return count, nil
	} else {
		data, err := readData(reader, size)
		if err != nil {
			return -1, err
		}

		switch typ {
		case uintegerType:
			handler.HandleInteger(id, int64(binary.BigEndian.Uint64(pad(data, 8))), info)
		case integerType:
			handler.HandleInteger(id, convertBytesToSignedInt(data), info)
		case binaryType:
			handler.HandleBinary(id, data, info)
		case stringType, utf8Type:
			handler.HandleString(id, string(unpadString(data)), info)
		case floatType:
			var value float64
			if size == 4 {
				value = float64(math.Float32frombits(binary.BigEndian.Uint32(data)))
			} else if size == 8 {
				value = math.Float64frombits(binary.BigEndian.Uint64(data))
			} else {
				return -1, fmt.Errorf("unexpected float size: %d", size)
			}
			handler.HandleFloat(id, value, info)
		case dateType:
			handler.HandleDate(id, baseDate.Add(time.Duration(convertBytesToSignedInt(data))), info)
		}
		return count, nil
	}
}

// Gives the human-readable name for the given element ID.
func NameForElementID(id ElementID) string {
	name, ok := elementNames[id]
	if !ok {
		return fmt.Sprintf("UNKNOWN:%x", id)
	}
	return name
}

////////////////////////////////////////////////////////////////////////////////
// Utility
////////////////////////////////////////////////////////////////////////////////

var baseDate = time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

// readData reads and returns size bytes from r.
// An error is returned if EOF is encountered before the requested bytes have been read.
func readData(r io.Reader, size int64) ([]byte, error) {
	// Use bytes.Buffer to avoid allocating the full size until needed:
	// https://github.com/remko/go-mkvparse/issues/4
	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, r, size); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

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
	if len(b) == size {
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
	_, err := reader.Read(b)
	if err != nil {
		return -1, -1, err
	}

	var mask byte
	var length int
	if ((b[0] & 0x80) >> 7) == 1 {
		length = 1
		mask = 0x7f
	} else if ((b[0] & 0x40) >> 6) == 1 {
		length = 2
		mask = 0x3f
	} else if ((b[0] & 0x20) >> 5) == 1 {
		length = 3
		mask = 0x1f
	} else if ((b[0] & 0x10) >> 4) == 1 {
		length = 4
		mask = 0xf
	} else if ((b[0] & 0x08) >> 3) == 1 {
		length = 5
		mask = 0x7
	} else if ((b[0] & 0x04) >> 2) == 1 {
		length = 6
		mask = 0x3
	} else if ((b[0] & 0x02) >> 1) == 1 {
		length = 7
		mask = 0x1
	} else if ((b[0] & 0x01) >> 0) == 1 {
		length = 8
		mask = 0x0
	}

	result := make([]byte, 8)
	if doMask {
		result[8-length] = b[0] & mask
	} else {
		result[8-length] = b[0]
	}
	_, err = reader.Read(result[8-length+1:])
	if err != nil {
		return -1, -1, err
	}
	return int64(binary.BigEndian.Uint64(result)), int64(length), nil
}
