package mkvparse

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

func readVarInt(reader io.Reader) (int64, int64, bool, error) {
	return readVarIntRaw(reader, true)
}

func readVarIntRaw(reader io.Reader, doMask bool) (int64, int64, bool, error) {
	b := make([]byte, 1)
	_, err := reader.Read(b)
	if err != nil {
		return -1, -1, false, err
	}

	var mask byte
	var allmask uint64
	var length int
	if ((b[0] & 0x80) >> 7) == 1 {
		length = 1
		mask = 0x7f
		allmask = 0x7f
	} else if ((b[0] & 0x40) >> 6) == 1 {
		length = 2
		mask = 0x3f
		allmask = 0x3fff
	} else if ((b[0] & 0x20) >> 5) == 1 {
		length = 3
		mask = 0x1f
		allmask = 0x1fffff
	} else if ((b[0] & 0x10) >> 4) == 1 {
		length = 4
		mask = 0xf
		allmask = 0x0fffffff
	} else if ((b[0] & 0x08) >> 3) == 1 {
		length = 5
		mask = 0x7
		allmask = 0x07ffffffff
	} else if ((b[0] & 0x04) >> 2) == 1 {
		length = 6
		mask = 0x3
		allmask = 0x03ffffffffff
	} else if ((b[0] & 0x02) >> 1) == 1 {
		length = 7
		mask = 0x1
		allmask = 0x01ffffffffffff
	} else if ((b[0] & 0x01) >> 0) == 1 {
		length = 8
		mask = 0x0
		allmask = 0x00ffffffffffffff
	} else {
		return -1, -1, false, fmt.Errorf("invalid varint length")
	}

	result := make([]byte, 8)
	if doMask {
		result[8-length] = b[0] & mask
	} else {
		result[8-length] = b[0]
	}
	_, err = reader.Read(result[8-length+1:])
	if err != nil {
		return -1, -1, false, err
	}

	uiresult := binary.BigEndian.Uint64(result)
	return int64(uiresult), int64(length), (uiresult & allmask) == (uint64(math.MaxUint64) & allmask), nil
}
