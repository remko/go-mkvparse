package mkvparse

import (
	"fmt"
	"log"
	"os"
)

func ExampleParseSections() {
	handler := myTitleHandler{}
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

type myTitleHandler struct {
	DefaultHandler
}

func (p *myTitleHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	switch id {
	case TitleElement:
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}
