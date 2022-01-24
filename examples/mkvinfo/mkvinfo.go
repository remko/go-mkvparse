// Prints all information of an MKV file
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/remko/go-mkvparse"
)

type MyParser struct {
	sawCluster bool
}

func (p *MyParser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	switch id {
	case mkvparse.CuesElement:
		fmt.Printf("%s- %s: <Skipping>\n", indent(info.Level), mkvparse.NameForElementID(id))
		return false, nil
	case mkvparse.ClusterElement:
		if !p.sawCluster {
			p.sawCluster = true
			fmt.Printf("%s- %s: <Skipping>\n", indent(info.Level), mkvparse.NameForElementID(id))
		}
		return false, nil
	default:
		fmt.Printf("%s- %s:\n", indent(info.Level), mkvparse.NameForElementID(id))
		return true, nil
	}
}

func (p *MyParser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	fmt.Printf("%s- %v: %q\n", indent(info.Level), mkvparse.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	fmt.Printf("%s- %v: %v\n", indent(info.Level), mkvparse.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	fmt.Printf("%s- %v: %v\n", indent(info.Level), mkvparse.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleDate(id mkvparse.ElementID, value time.Time, info mkvparse.ElementInfo) error {
	fmt.Printf("%s- %v: %v\n", indent(info.Level), mkvparse.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.SeekIDElement:
		fmt.Printf("%s- %v: %x\n", indent(info.Level), mkvparse.NameForElementID(id), value)
	default:
		fmt.Printf("%s- %v: <binary>\n", indent(info.Level), mkvparse.NameForElementID(id))
	}
	return nil
}

func main() {
	handler := MyParser{}
	err := mkvparse.ParsePath(os.Args[1], &handler)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}

func indent(n int) string {
	return strings.Repeat("  ", n)
}
