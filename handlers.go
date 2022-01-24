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
