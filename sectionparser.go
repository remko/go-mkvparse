package mkvparse

import (
	"os"
	"time"
)

type sectionsHandler struct {
	DefaultHandler

	sections        map[ElementID]bool
	delegateHandler Handler
}

func (p *sectionsHandler) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	_, ok := p.sections[id]
	if ok {
		return p.delegateHandler.HandleMasterBegin(id, info)
	}

	return false, nil
}

func (p *sectionsHandler) HandleMasterEnd(id ElementID, info ElementInfo) error {
	_, ok := p.sections[id]
	if ok {
		return p.delegateHandler.HandleMasterEnd(id, info)
	}

	return nil
}

func (p *sectionsHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	return p.delegateHandler.HandleString(id, value, info)
}

func (p *sectionsHandler) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	return p.delegateHandler.HandleInteger(id, value, info)
}

func (p *sectionsHandler) HandleFloat(id ElementID, value float64, info ElementInfo) error {
	return p.delegateHandler.HandleFloat(id, value, info)
}

func (p *sectionsHandler) HandleDate(id ElementID, value time.Time, info ElementInfo) error {
	return p.delegateHandler.HandleDate(id, value, info)
}

func (p *sectionsHandler) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
	return p.delegateHandler.HandleBinary(id, value, info)
}

// Parses only the given sections of `file`.
//
// When present, uses the seek index to avoid having to parse the entire file
func ParseSections(file *os.File, sections []ElementID, handler Handler) error {
	sectionsHandler := sectionsHandler{
		sections:        map[ElementID]bool{},
		delegateHandler: handler,
	}

	tmp := map[ElementID]bool{}

	for _, section := range sections {
		tmp[section] = true
	}

	for k, _ := range tmp {
		sectionsHandler.sections[k] = true

		parents := ParentsForElementID(k)

		for _, p := range parents {
			sectionsHandler.sections[p] = true
		}
	}

	return Parse(file, &sectionsHandler)
}
