package mkvparse

import (
	"fmt"
	"os"
)

func ExampleHandlerChain() {
	file, err := os.Open("testdata/example-cover.mkv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	coverh := CoverHandler{}
	titleh := TitleHandler{}
	if err := ParseSections(file, NewHandlerChain(&coverh, &titleh), InfoElement, AttachmentsElement); err != nil {
		panic(err)
	}
	fmt.Printf("parsed cover: %s (%d bytes)\n", coverh.MIMEType, len(coverh.Data))

	// Output:
	// Title: Awesome Movie
	// parsed cover: image/jpeg (41363 bytes)
}

type TitleHandler struct {
	DefaultHandler
}

func (p *TitleHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	switch id {
	case TitleElement:
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}
