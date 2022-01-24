package mkvparse

import (
	"fmt"
	"log"
)

type MyHandler struct {
	DefaultHandler
}

func (p *MyHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	switch id {
	case TitleElement:
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}

func Example() {
	handler := MyHandler{}
	if err := ParsePath("testdata/example.mkv", &handler); err != nil {
		log.Fatalf("%v", err)
	}
	// Output:
	// Title: Awesome Movie
}
