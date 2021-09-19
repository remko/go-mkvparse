package mkvparse

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
	"time"
)

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
	events      []ParseEvent
	skipDescend bool
}

func (p *ParseHandler) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	p.events = append(p.events, ParseEvent{id, info, MasterBeginEvent{}})
	return !p.skipDescend, nil
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
	data   []byte
	events []ParseEvent
	fail   bool // error is expected
}

func TestParseElement(t *testing.T) {
	tests := map[string]ParseTest{
		"time before millenium": {
			[]byte{0x44, 0x61, 0x88, 0xf6, 0xd3, 0xc2, 0xb9, 0x1b, 0xee, 0x28, 0x00},
			[]ParseEvent{{
				DateUTCElement,
				ElementInfo{
					Offset: 3,
					Size:   8,
					Level:  0,
				},
				time.Date(1980, time.January, 21, 21, 03, 0, 0, time.UTC),
			}},
			false,
		},
		"master": {
			data: []byte{
				0x1F, 0x43, 0xB6, 0x75, 0x80 | 0x3,
				0xE7, 0x80 | 0x1, 0x3,
			},
			events: []ParseEvent{
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   3,
						Level:  0,
					},
					MasterBeginEvent{},
				},
				{
					TimecodeElement,
					ElementInfo{
						Offset: 7,
						Size:   1,
						Level:  1,
					},
					int64(0x3),
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   3,
						Level:  0,
					},
					MasterEndEvent{},
				},
			},
		},
		"unknown type": {
			data: []byte{
				0x1A, 0x45, 0xDF, 0xA4, 0x80 | 0x6,
				0xE7, 0x80 | 0x1, 0x3,
				0xE7, 0x80 | 0x1, 0x3,
			},
			events: nil,
		},
		"invalid integer size": {
			data: []byte{
				0xE7, 0x80 | 0xa, 0x10, 0x00, 0x00, 0x02, 0x10, 0x00, 0x00, 0x02, 0x40, 0x02,
			},
			events: nil,
			fail:   true,
		},

		"master (unknown size)": {
			data: []byte{
				0x1F, 0x43, 0xB6, 0x75, 0xFF,
				0xE7, 0x80 | 0x1, 0x3,
				0xE7, 0x80 | 0x1, 0x3,
				0xE7, 0x80 | 0x1, 0x3,
				0x1F, 0x43, 0xB6, 0x75, 0x80 | 0x3,
				0xE7, 0x80 | 0x1, 0x3,
			},
			events: []ParseEvent{
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   -1,
						Level:  0,
					},
					MasterBeginEvent{},
				},
				{
					TimecodeElement,
					ElementInfo{
						Offset: 7,
						Size:   1,
						Level:  1,
					},
					int64(0x3),
				},
				{
					TimecodeElement,
					ElementInfo{
						Offset: 10,
						Size:   1,
						Level:  1,
					},
					int64(0x3),
				},
				{
					TimecodeElement,
					ElementInfo{
						Offset: 13,
						Size:   1,
						Level:  1,
					},
					int64(0x3),
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   -1,
						Level:  0,
					},
					MasterEndEvent{},
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 20,
						Size:   3,
						Level:  0,
					},
					MasterBeginEvent{},
				},
				{
					TimecodeElement,
					ElementInfo{
						Offset: 22,
						Size:   1,
						Level:  1,
					},
					int64(0x3),
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 20,
						Size:   3,
						Level:  0,
					},
					MasterEndEvent{},
				},
			},

			// TODO: Test unknown size in unknown size (e.g.  \Segment\Cluster(unknown)\BlockGroup(unknown)\BlockDuration)
		},

		// Avoid panicking with a too-large slice allocation when an element claims a
		// very large size: https://github.com/remko/go-mkvparse/issues/4
		"excessive size": {
			[]byte{0xa3, 0x01, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x02, 0x03, 0x04},
			[]ParseEvent{},
			true,
		},
	}
	for name, test := range tests {
		name := name
		test := test
		t.Run(name, func(t *testing.T) {
			reader := bytes.NewReader(test.data)
			handler := ParseHandler{}
			count, nextID, _, err := parseElement(reader, 0, 0, -1, &handler)
			if nextID != -1 {
				t.Fatalf("unexpected next ID")
			}
			if test.fail {
				if err == nil {
					t.Fatal("Unexpectedly succeeded")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if count != int64(len(test.data)) {
					t.Fatalf("Invalid #bytes read: %d != %d. Data: %v", count, len(test.data), test.data)
				}
				if !reflect.DeepEqual(test.events, handler.events) {
					t.Fatalf("Invalid events: %#v != %#v", test.events, handler.events)
				}
			}
		})
	}
}

func TestParseMaster_Skips(t *testing.T) {
	tests := []struct {
		data           []byte
		expectedEvents []ParseEvent
	}{
		{
			data: []byte{
				0x1F, 0x43, 0xB6, 0x75, 0x80 | 0x3,
				0xE7, 0x80 | 0x1, 0x3,
			},
			expectedEvents: []ParseEvent{
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   3,
						Level:  0,
					},
					MasterBeginEvent{},
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   3,
						Level:  0,
					},
					MasterEndEvent{},
				},
			},
		},

		// Unknown size
		{
			data: []byte{
				0x1F, 0x43, 0xB6, 0x75, 0xFF,
				0xE7, 0x80 | 0x1, 0x3,
				0xE7, 0x80 | 0x1, 0x3,
				0xE7, 0x80 | 0x1, 0x3,
				0x1F, 0x43, 0xB6, 0x75, 0x80 | 0x3,
				0xE7, 0x80 | 0x1, 0x3,
			},
			expectedEvents: []ParseEvent{
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   -1,
						Level:  0,
					},
					MasterBeginEvent{},
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 5,
						Size:   -1,
						Level:  0,
					},
					MasterEndEvent{},
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 20,
						Size:   3,
						Level:  0,
					},
					MasterBeginEvent{},
				},
				{
					ClusterElement,
					ElementInfo{
						Offset: 20,
						Size:   3,
						Level:  0,
					},
					MasterEndEvent{},
				},
			},

			// TODO: Test unknown size in unknown size (e.g.  \Segment\Cluster(unknown)\BlockGroup(unknown)\BlockDuration)
		},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test.data), func(t *testing.T) {
			readers := [](func([]byte) io.Reader){
				func(b []byte) io.Reader { return bytes.NewReader(b) },
				func(b []byte) io.Reader { return bytes.NewBuffer(b) },
			}

			for _, reader := range readers {
				reader := reader(test.data)
				handler := ParseHandler{skipDescend: true}
				count, nextID, _, err := parseElement(reader, 0, 0, -1, &handler)
				if nextID != -1 {
					t.Fatalf("unexpected next ID")
				}
				if err != nil {
					t.Errorf("%v", err)
					continue
				}
				if count != int64(len(test.data)) {
					t.Errorf("Invalid #bytes read: %d != %d. Data: %v", count, len(test.data), test.data)
				}
				if !reflect.DeepEqual(test.expectedEvents, handler.events) {
					t.Errorf("Invalid events: %#v != %#v", test.expectedEvents, handler.events)
				}
			}
		})
	}
}

func TestParseMaster_SkipsWithInsufficientData(t *testing.T) {
	data := []byte{0x1F, 0x43, 0xB6, 0x75, 0x80 | 0x3, 0xE7}
	reader := bytes.NewBuffer(data)
	handler := ParseHandler{skipDescend: true}
	_, nextID, _, err := parseElement(reader, 0, 0, -1, &handler)
	if err == nil {
		t.Errorf("unexpected success")
	}
	if nextID != -1 {
		t.Fatalf("unexpected next ID")
	}
}
