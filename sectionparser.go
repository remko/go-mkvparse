package mkvparse

import (
	"encoding/binary"
	"errors"
	"io"
	"time"
)

var errAbort = errors.New("abort")

type sectionsHandler struct {
	sections            map[ElementID]bool
	seenSections        map[int64]bool
	segmentInfo         ElementInfo
	index               map[ElementID]int64
	delegateHandler     Handler
	inSeekHead          bool
	currentSeekPosition int64
	currentSeekID       ElementID
}

func (p *sectionsHandler) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	if id == SegmentElement {
		p.segmentInfo = info
		return p.delegateHandler.HandleMasterBegin(id, info)
	} else if info.Level <= 1 {
		if id == SeekHeadElement {
			p.index = make(map[ElementID]int64)
			p.inSeekHead = true
			return true, nil
		} else if p.sections[id] && !p.seenSections[info.Offset] {
			return p.delegateHandler.HandleMasterBegin(id, info)
		} else {
			return false, nil
		}
	} else if p.inSeekHead {
		return true, nil
	} else {
		return p.delegateHandler.HandleMasterBegin(id, info)
	}
}

func (p *sectionsHandler) HandleMasterEnd(id ElementID, info ElementInfo) error {
	if id == SegmentElement {
		return p.delegateHandler.HandleMasterEnd(id, info)
	} else if info.Level <= 1 {
		if id == SeekHeadElement {
			return errAbort
		} else if p.sections[id] && !p.seenSections[info.Offset] {
			p.seenSections[info.Offset] = true
			return p.delegateHandler.HandleMasterEnd(id, info)
		}
	} else if p.inSeekHead {
		if id == SeekElement {
			p.index[p.currentSeekID] = p.segmentInfo.Offset + p.currentSeekPosition
		}
		return nil
	} else {
		return p.delegateHandler.HandleMasterEnd(id, info)
	}
	return nil
}

func (p *sectionsHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	if p.inSeekHead {
		return nil
	} else {
		return p.delegateHandler.HandleString(id, value, info)
	}
}

func (p *sectionsHandler) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	if p.inSeekHead {
		if id == SeekPositionElement {
			p.currentSeekPosition = value
		}
		return nil
	} else {
		return p.delegateHandler.HandleInteger(id, value, info)
	}
}

func (p *sectionsHandler) HandleFloat(id ElementID, value float64, info ElementInfo) error {
	if p.inSeekHead {
		return nil
	} else {
		return p.delegateHandler.HandleFloat(id, value, info)
	}
}

func (p *sectionsHandler) HandleDate(id ElementID, value time.Time, info ElementInfo) error {
	if p.inSeekHead {
		return nil
	} else {
		return p.delegateHandler.HandleDate(id, value, info)
	}
}

func (p *sectionsHandler) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
	if p.inSeekHead {
		if id == SeekIDElement {
			p.currentSeekID = ElementID(binary.BigEndian.Uint64(pad(value, 8)))
		}
		return nil
	} else {
		return p.delegateHandler.HandleBinary(id, value, info)
	}
}

// Parses only the given sections of `file`.
//
// When present, uses the seek index to avoid having to parse the entire file
func ParseSections(file io.ReadSeeker, handler Handler, sections ...ElementID) error {
	sectionsHandler := sectionsHandler{
		sections:        make(map[ElementID]bool),
		seenSections:    make(map[int64]bool),
		delegateHandler: handler,
	}
	for _, section := range sections {
		sectionsHandler.sections[section] = true
	}

	// First pass
	err := Parse(file, &sectionsHandler)
	if err == errAbort {
		// Second pass
		for _, section := range sections {
			sectionOffset, ok := sectionsHandler.index[section]
			if ok && !sectionsHandler.seenSections[sectionOffset] {
				if _, err := file.Seek(sectionOffset, io.SeekStart); err != nil {
					return err
				}
				if _, _, _, err = parseElement(file, sectionOffset, 1, -1, handler); err != nil {
					return err
				}
				sectionsHandler.seenSections[sectionOffset] = true
			}
		}
		return handler.HandleMasterEnd(SegmentElement, sectionsHandler.segmentInfo)
	} else if err != nil {
		return err
	}
	return nil
}
