package mkvparse

import (
	"fmt"
	"os"

	"github.com/remko/go-mkvparse"
)

func ExampleCoverHandler() {
	file, err := os.Open("../example-cover.mkv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	handler := CoverHandler{}
	err = mkvparse.ParseSections(file, &handler, mkvparse.AttachmentsElement)
	if err != nil {
		panic(err)
	}

	fmt.Printf("parsed cover: %s (%d bytes)\n", handler.MIMEType, len(handler.Data))

	// Output:
	// parsed cover: image/jpeg (41363 bytes)
}
