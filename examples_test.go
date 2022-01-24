package mkvparse

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type TitleParser struct {
	DefaultHandler
}

////////////////////////////////////////////////////////////////////////////////

func (p *TitleParser) HandleString(id ElementID, value string, info ElementInfo) error {
	switch id {
	case TitleElement:
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}
func ExampleNameForElementID() {
	fmt.Println(NameForElementID(InfoElement))
	// Output:
	// Info
}

////////////////////////////////////////////////////////////////////////////////

func ExampleParseSections() {
	handler := TitleParser{}
	file, err := os.Open("testdata/example.mkv")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	if err = ParseSections(file, &handler, InfoElement); err != nil {
		log.Fatalf("%v", err)
	}
	// Output:
	// Title: Awesome Movie
}

////////////////////////////////////////////////////////////////////////////////

func Example() {
	handler := TitleParser{}
	if err := ParsePath("testdata/example.mkv", &handler); err != nil {
		log.Fatalf("%v", err)
	}
	// Output:
	// Title: Awesome Movie
}

////////////////////////////////////////////////////////////////////////////////

func ExampleCoverHandler() {
	file, err := os.Open("testdata/example-cover.mkv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	handler := CoverHandler{}
	err = ParseSections(file, &handler, AttachmentsElement)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed cover: %s (%d bytes)\n", handler.MIMEType, len(handler.Data))

	// Output:
	// parsed cover: image/jpeg (41363 bytes)
}

func ExampleParseCover() {
	data, typ, err := ParseCover("testdata/example-cover.mkv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed cover: %s (%d bytes)\n", typ, len(data))

	// Output:
	// parsed cover: image/jpeg (41363 bytes)
}

func ExampleParseCover_image() {
	data, typ, err := ParseCover("testdata/example-cover.mkv")
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
