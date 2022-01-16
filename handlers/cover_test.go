package mkvparse

import (
	"fmt"
)

func ExampleParseCover() {
	data, typ, err := ParseCover("../example-cover.mkv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed cover: %s (%d bytes)\n", typ, len(data))

	// Output:
	// parsed cover: image/jpeg (41363 bytes)
}
