package mkvparse

import (
	"os"
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

func main() {}
