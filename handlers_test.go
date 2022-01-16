package mkvparse

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
)

func ExampleParseCover() {
	data, typ, err := ParseCover("example-cover.mkv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed cover: %s (%d bytes)\n", typ, len(data))

	// Output:
	// parsed cover: image/jpeg (41363 bytes)
}

func ExampleParseCover_Image() {
	data, typ, err := ParseCover("example-cover.mkv")
	if err != nil {
		log.Panic(err)
	}
	if data == nil {
		log.Panic("no cover")
		return
	}
	var img image.Image
	switch typ {
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(data))
	case "image/png":
		img, err = png.Decode(bytes.NewReader(data))
	default:
		log.Panicf("unknown MIME type: %s", typ)
	}
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("parsed cover image: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

	// Output:
	// parsed cover image: 265x377
}
