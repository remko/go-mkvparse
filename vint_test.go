package mkvparse

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadVarInt(t *testing.T) {
	tests := []struct {
		in      []byte
		result  int64
		size    int64
		isAll1  bool
		isError bool
	}{
		// All 1 encodings
		{in: []byte{0xff}, result: 127, size: 1, isAll1: true, isError: false},
		{in: []byte{0x40, 0x7f}, result: 127, size: 2, isAll1: false, isError: false},
		{in: []byte{0x7f, 0xff}, result: 16383, size: 2, isAll1: true, isError: false},

		// Zero length
		{in: []byte{0x0}, result: 0, size: 0, isAll1: false, isError: true},

		// Different encodings of '2'
		{in: []byte{0x82}, result: 2, size: 1, isAll1: false, isError: false},
		{in: []byte{0x40, 0x02}, result: 2, size: 2, isAll1: false, isError: false},
		{in: []byte{0x20, 0x00, 0x02}, result: 2, size: 3, isAll1: false, isError: false},
		{in: []byte{0x10, 0x00, 0x00, 0x02}, result: 2, size: 4, isAll1: false, isError: false},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			in := make([]byte, len(test.in))
			copy(in, test.in)
			in = append(in, 0xde, 0xad, 0xbe, 0xef)
			reader := bytes.NewReader(in)
			result, count, isAll1, err := readVarInt(reader)
			if test.isError {
				if err == nil {
					t.Fatalf("unexpected succes")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if count != test.size {
					t.Fatalf("unexpected count: %d != %d", count, test.size)
				}
				if result != test.result {
					t.Fatalf("unexpected result: %d != %d", result, test.result)
				}
				if isAll1 != test.isAll1 {
					t.Fatalf("unexpected all1 %v != %v", isAll1, test.isAll1)
				}
			}

		})
	}
}
