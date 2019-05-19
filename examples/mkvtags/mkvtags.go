////////////////////////////////////////////////////////////
// Prints tags of an MKV file
////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/remko/go-mkvparse"
	"os"
	"sort"
	"time"
)

type MyParser struct {
	currentTagGlobal bool
	currentTagName   *string
	currentTagValue  *string
	title            *string
	tags             map[string]string
}

func (p *MyParser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	if id == mkvparse.TagElement {
		p.currentTagGlobal = true
	} else if id == mkvparse.SimpleTagElement {
		p.currentTagName = nil
		p.currentTagValue = nil
	}
	return true, nil
}

func (p *MyParser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	if id == mkvparse.SimpleTagElement && p.currentTagGlobal && p.currentTagName != nil && p.currentTagValue != nil {
		p.tags[*p.currentTagName] = *p.currentTagValue
	}
	return nil
}

func (p *MyParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	if id == mkvparse.TagNameElement {
		p.currentTagName = &value
	} else if id == mkvparse.TagStringElement {
		p.currentTagValue = &value
	} else if id == mkvparse.TitleElement {
		p.title = &value
	}
	return nil
}

func (p *MyParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	if (id == mkvparse.TagTrackUIDElement || id == mkvparse.TagEditionUIDElement || id == mkvparse.TagChapterUIDElement || id == mkvparse.TagAttachmentUIDElement) && value != 0 {
		p.currentTagGlobal = false
	}
	return nil
}

func (p *MyParser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleDate(id mkvparse.ElementID, value time.Time, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	return nil
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	defer file.Close()
	handler := MyParser{
		tags: make(map[string]string),
	}
	err = mkvparse.ParseSections(file, &handler, mkvparse.InfoElement, mkvparse.TagsElement)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	// Print (sorted) tags
	if handler.title != nil {
		fmt.Printf("- title: %q\n", *handler.title)
	}
	var tagNames []string
	for tagName := range handler.tags {
		tagNames = append(tagNames, tagName)
	}
	sort.Strings(tagNames)
	for _, tagName := range tagNames {
		fmt.Printf("- %s: %q\n", tagName, handler.tags[tagName])
	}
}
