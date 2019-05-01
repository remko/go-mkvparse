package mkvparse

import (
	"fmt"
	"log"
	"os"
	"time"
)

type myTitleParser struct {
	DefaultHandler
}

func (p *myTitleParser) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	return true, nil
}

func (p *myTitleParser) HandleMasterEnd(id ElementID, info ElementInfo) error {
	return nil
}

func (p *myTitleParser) HandleString(id ElementID, value string, info ElementInfo) error {
	if id == TitleElement {
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}

func (p *myTitleParser) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	return nil
}

func (p *myTitleParser) HandleFloat(id ElementID, value float64, info ElementInfo) error {
	return nil
}

func (p *myTitleParser) HandleDate(id ElementID, value time.Time, info ElementInfo) error {
	return nil
}

func (p *myTitleParser) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
	return nil
}

func ExampleParseSections() {
	file, err := os.Open("example.mkv")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	handler := myTitleParser{}

	if err = ParseSections(file, []ElementID{InfoElement}, &handler); err != nil {
		log.Fatalf("%v", err)
	}

	// Output:
	// Title: Awesome Movie
}
