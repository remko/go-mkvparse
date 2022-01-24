package mkvparse

import (
	"os"
	"time"
)

////////////////////////////////////////////////////////////////////////////////

// Handler to extract cover image data
//
// Needs the section parser to handle `AttachmentsElement`
type CoverHandler struct {
	DefaultHandler

	currentAttachmentData     []byte
	currentAttachmentFileName string
	currentAttachmentMIMEType string

	Data     []byte
	MIMEType string
}

func (p *CoverHandler) HandleMasterEnd(id ElementID, info ElementInfo) error {
	if id == AttachedFileElement && (p.currentAttachmentFileName == "cover.jpg" || p.currentAttachmentFileName == "cover.png") {
		p.Data = p.currentAttachmentData
		p.MIMEType = p.currentAttachmentMIMEType
	}
	return nil
}

func (p *CoverHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	if id == FileNameElement {
		p.currentAttachmentFileName = value
	} else if id == FileMimeTypeElement {
		p.currentAttachmentMIMEType = value
	}
	return nil
}

func (p *CoverHandler) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
	if id == FileDataElement {
		p.currentAttachmentData = value
	}
	return nil
}

func ParseCover(path string) ([]byte, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	handler := CoverHandler{}
	err = ParseSections(file, &handler, AttachmentsElement)
	if err != nil {
		return nil, "", err
	}

	return handler.Data, handler.MIMEType, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Handler that chains multiple handlers.
//
// All handlers are called in sequence. If one of the handers requests to descend, the parser will descend.
type HandlerChain struct {
	Handlers []Handler
}

func (c *HandlerChain) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	descend := false
	for _, h := range c.Handlers {
		d, err := h.HandleMasterBegin(id, info)
		if err != nil {
			return descend, err
		}
		descend = descend || d
	}
	return descend, nil
}

func (c *HandlerChain) HandleMasterEnd(id ElementID, info ElementInfo) error {
	for _, h := range c.Handlers {
		if err := h.HandleMasterEnd(id, info); err != nil {
			return err
		}
	}
	return nil
}

func (c *HandlerChain) HandleString(id ElementID, value string, info ElementInfo) error {
	for _, h := range c.Handlers {
		if err := h.HandleString(id, value, info); err != nil {
			return err
		}
	}
	return nil
}

func (c *HandlerChain) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	for _, h := range c.Handlers {
		if err := h.HandleInteger(id, value, info); err != nil {
			return err
		}
	}
	return nil
}

func (c *HandlerChain) HandleFloat(id ElementID, value float64, info ElementInfo) error {
	for _, h := range c.Handlers {
		if err := h.HandleFloat(id, value, info); err != nil {
			return err
		}
	}
	return nil
}

func (c *HandlerChain) HandleDate(id ElementID, value time.Time, info ElementInfo) error {
	for _, h := range c.Handlers {
		if err := h.HandleDate(id, value, info); err != nil {
			return err
		}
	}
	return nil
}

func (c *HandlerChain) HandleBinary(id ElementID, value []byte, info ElementInfo) error {
	for _, h := range c.Handlers {
		if err := h.HandleBinary(id, value, info); err != nil {
			return err
		}
	}
	return nil
}

// Creates a new handler that chains `handlers`
func NewHandlerChain(handlers ...Handler) *HandlerChain {
	return &HandlerChain{
		Handlers: handlers,
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Handler that collects tags.
//
// Needs the section parser to handle `TagsElement`.
type TagsHandler struct {
	DefaultHandler

	currentTagTrackUIDs      []int64
	currentTagChapterUIDs    []int64
	currentTagAttachmentUIDs []int64
	currentTagEditionUIDs    []int64

	currentTags     map[string]string
	currentTagName  *string
	currentTagValue *string

	tags           map[string]string
	trackTags      map[int64]map[string]string
	chapterTags    map[int64]map[string]string
	attachmentTags map[int64]map[string]string
	editionTags    map[int64]map[string]string
}

func NewTagsHandler() *TagsHandler {
	return &TagsHandler{
		tags:           map[string]string{},
		trackTags:      map[int64]map[string]string{},
		chapterTags:    map[int64]map[string]string{},
		attachmentTags: map[int64]map[string]string{},
		editionTags:    map[int64]map[string]string{},
	}
}

// Retrieves all global tags
func (h *TagsHandler) Tags() map[string]string {
	return h.tags
}

// Retrieves tags for track with UID `uid`.
//
// Returns nil if no tags were encountered for this track.
func (h *TagsHandler) TrackTags(uid int64) map[string]string {
	return h.trackTags[uid]
}

// Retrieves tags for chapter with UID `uid`.
//
// Returns nil if no tags were encountered for this chapter.
func (h *TagsHandler) ChapterTags(uid int64) map[string]string {
	return h.chapterTags[uid]
}

// Retrieves tags for attachment with UID `uid`.
//
// Returns nil if no tags were encountered for this attachment.
func (h *TagsHandler) AttachmentTags(uid int64) map[string]string {
	return h.attachmentTags[uid]
}

// Retrieves tags for edition with UID `uid`.
//
// Returns nil if no tags were encountered for this edition.
func (h *TagsHandler) EditionTags(uid int64) map[string]string {
	return h.editionTags[uid]
}

func (p *TagsHandler) HandleMasterBegin(id ElementID, info ElementInfo) (bool, error) {
	switch id {
	case TagElement:
		p.resetTagState()
		p.currentTags = map[string]string{}

	case SimpleTagElement:
		p.resetSimpleTagState()
	}
	return true, nil
}

func (p *TagsHandler) HandleMasterEnd(id ElementID, info ElementInfo) error {
	switch id {
	case TagElement:
		if len(p.currentTagAttachmentUIDs) == 0 && len(p.currentTagChapterUIDs) == 0 && len(p.currentTagAttachmentUIDs) == 0 && len(p.currentTagEditionUIDs) == 0 {
			for k, v := range p.currentTags {
				p.tags[k] = v
			}
		} else {
			for _, uid := range p.currentTagTrackUIDs {
				t := p.trackTags[uid]
				if t == nil {
					p.trackTags[uid] = p.currentTags
				} else {
					for k, v := range p.currentTags {
						t[k] = v
					}
				}
			}
			for _, uid := range p.currentTagChapterUIDs {
				t := p.chapterTags[uid]
				if t == nil {
					p.chapterTags[uid] = p.currentTags
				} else {
					for k, v := range p.currentTags {
						t[k] = v
					}
				}
			}
			for _, uid := range p.currentTagAttachmentUIDs {
				t := p.attachmentTags[uid]
				if t == nil {
					p.attachmentTags[uid] = p.currentTags
				} else {
					for k, v := range p.currentTags {
						t[k] = v
					}
				}
			}
			for _, uid := range p.currentTagEditionUIDs {
				t := p.editionTags[uid]
				if t == nil {
					p.editionTags[uid] = p.currentTags
				} else {
					for k, v := range p.currentTags {
						t[k] = v
					}
				}
			}
		}
		p.resetTagState()

	case SimpleTagElement:
		if p.currentTagName != nil && p.currentTagValue != nil {
			p.currentTags[*p.currentTagName] = *p.currentTagValue
		}
		p.resetSimpleTagState()
	}
	return nil
}

func (p *TagsHandler) HandleString(id ElementID, value string, info ElementInfo) error {
	switch id {
	case TagNameElement:
		p.currentTagName = &value
	case TagStringElement:
		p.currentTagValue = &value
	}
	return nil
}

func (p *TagsHandler) HandleInteger(id ElementID, value int64, info ElementInfo) error {
	switch id {
	case TagTrackUIDElement:
		p.currentTagTrackUIDs = append(p.currentTagTrackUIDs, value)
	case TagChapterUIDElement:
		p.currentTagChapterUIDs = append(p.currentTagChapterUIDs, value)
	case TagAttachmentUIDElement:
		p.currentTagAttachmentUIDs = append(p.currentTagAttachmentUIDs, value)
	case TagEditionUIDElement:
		p.currentTagEditionUIDs = append(p.currentTagEditionUIDs, value)
	}
	return nil
}

func (p *TagsHandler) resetTagState() {
	p.currentTags = nil
	p.currentTagTrackUIDs = nil
	p.currentTagChapterUIDs = nil
	p.currentTagAttachmentUIDs = nil
	p.currentTagEditionUIDs = nil
}

func (p *TagsHandler) resetSimpleTagState() {
	p.currentTagName = nil
	p.currentTagValue = nil
}
