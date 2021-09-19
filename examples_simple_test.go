package mkvparse

import (
	"fmt"
	"log"
)

type MyParser struct {
	DefaultHandler
}

func (p *MyParser) HandleString(id ElementID, value string, info ElementInfo) error {
	if id == TitleElement {
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}

func Example() {
	handler := MyParser{}
	if err := ParsePath("example.mkv", &handler); err != nil {
		log.Fatalf("%v", err)
	}
	// Output:
	// Title: Awesome Movie
}
