package mkvparse

import (
	"fmt"
	"log"
	"os"
)

type myTitleParser struct {
	DefaultHandler
}

func (p *myTitleParser) HandleString(id ElementID, value string, info ElementInfo) error {
	if id == TitleElement {
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}

func ExampleParseSections() {
	handler := myTitleParser{}
	file, err := os.Open("example.mkv")
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
