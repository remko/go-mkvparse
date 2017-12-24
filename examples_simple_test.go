package mkvparse

import (
	"fmt"
	"log"
	"time"
)

type MyParser struct {
}

func (p *MyParser) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	return true, nil
}

func (p *MyParser) HandleMasterEnd(id ElementID, info ElementInfo) error {
	return nil
}

func (p *MyParser) HandleString(id ElementID, value string, info ElementInfo) error {
	if id == TitleElement {
		fmt.Printf("%s: %v\n", NameForElementID(id), value)
	}
	return nil
}

func (p *MyParser) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	return nil
}

func (p *MyParser) HandleFloat(id ElementID, value float64, info ElementInfo) error {
	return nil
}

func (p *MyParser) HandleDate(id ElementID, value time.Time, info ElementInfo) error {
	return nil
}

func (p *MyParser) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
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
