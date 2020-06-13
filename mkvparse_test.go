package mkvparse

import (
	"bytes"
	"testing"
	"time"
)

func TestReadVarInt_2Encodings(t *testing.T) {
	testEncodings := [][]byte{
		{0x82},
		{0x40, 0x02},
		{0x20, 0x00, 0x02},
		{0x10, 0x00, 0x00, 0x02},
	}
	for _, encoding := range testEncodings {
		encoding = append(encoding, 0xde, 0xad, 0xbe, 0xef)
		reader := bytes.NewReader(encoding)
		result, count, err := readVarInt(reader)
		if err != nil {
			t.Errorf("%x: %v", encoding, err)
		}
		if count != int64(len(encoding))-4 {
			t.Errorf("%x: %d != %d", encoding, count, len(encoding)-4)
		}
		if result != 0x2 {
			t.Errorf("%x: %d != %d", encoding, result, 10)
		}
	}
}

func TestReadElementID(t *testing.T) {
	testIDs := map[ElementID][]byte{
		TimecodeElement:        {0xE7},
		EBMLVersionElement:     {0x42, 0x86},
		DefaultDurationElement: {0x23, 0xE3, 0x83},
		EBMLElement:            {0x1A, 0x45, 0xDF, 0xA3},
	}
	for id, encoding := range testIDs {
		encoding = append(encoding, 0xde, 0xad, 0xbe, 0xef)
		reader := bytes.NewReader(encoding)
		result, count, err := readElementID(reader)
		if err != nil {
			t.Errorf("%x: %v", encoding, err)
		}
		if count != int64(len(encoding))-4 {
			t.Errorf("%x: %d != %d", encoding, count, len(encoding)-4)
		}
		if result != id {
			t.Errorf("%x: %x != %x", encoding, result, id)
		}
	}
}

//////////////////////////////////////////////////////////////////////

type ParseEvent struct {
	id    ElementID
	info  ElementInfo
	value interface{}
}

type MasterBeginEvent struct{}
type MasterEndEvent struct{}

type ParseHandler struct {
	events []ParseEvent
}

func (p *ParseHandler) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	p.events = append(p.events, ParseEvent{id, info, MasterBeginEvent{}})
	return true, nil
}

func (p *ParseHandler) HandleMasterEnd(id ElementID, info ElementInfo) error {
	p.events = append(p.events, ParseEvent{id, info, MasterEndEvent{}})
	return nil
}

func (p *ParseHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	p.events = append(p.events, ParseEvent{id, info, value})
	return nil
}

func (p *ParseHandler) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	p.events = append(p.events, ParseEvent{id, info, value})
	return nil
}

func (p *ParseHandler) HandleFloat(id ElementID, value float64, info ElementInfo) error {
	p.events = append(p.events, ParseEvent{id, info, value})
	return nil
}

func (p *ParseHandler) HandleDate(id ElementID, value time.Time, info ElementInfo) error {
	p.events = append(p.events, ParseEvent{id, info, value})
	return nil
}

func (p *ParseHandler) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
	p.events = append(p.events, ParseEvent{id, info, value})
	return nil
}

type ParseTest struct {
	data  []byte
	event ParseEvent
	fail  bool // error is expected
}

func TestParseElement(t *testing.T) {
	tests := map[string]ParseTest{
		"time before millenium": {
			[]byte{0x44, 0x61, 0x88, 0xf6, 0xd3, 0xc2, 0xb9, 0x1b, 0xee, 0x28, 0x00},
			ParseEvent{
				DateUTCElement,
				ElementInfo{
					Offset: 3,
					Size:   8,
					Level:  0,
				},
				time.Date(1980, time.January, 21, 21, 03, 0, 0, time.UTC),
			},
			false,
		},
		// Avoid panicking with a too-large slice allocation when an element claims a
		// very large size: https://github.com/remko/go-mkvparse/issues/4
		"excessive size": {
			[]byte{0xa3, 0x01, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x02, 0x03, 0x04},
			ParseEvent{},
			true,
		},
	}
	for name, test := range tests {
		reader := bytes.NewReader(test.data)
		handler := ParseHandler{}
		count, err := parseElement(reader, 0, 0, &handler)
		if test.fail {
			if err == nil {
				t.Errorf("%s: Unexpectedly succeeded", name)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %v", name, err)
			}
			if count != int64(len(test.data)) {
				t.Errorf("%s: %d != %d. Data: %v", name, count, len(test.data), test.data)
			}
			if len(handler.events) != 1 {
				t.Errorf("%s: Invalid event count: %d", name, len(handler.events))
			}
			if test.event != handler.events[0] {
				t.Errorf("%s: Invalid event: %v != %v", name, test.event, handler.events[0])
			}
		}
	}
}
