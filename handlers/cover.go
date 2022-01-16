package mkvparse

import (
	"github.com/remko/go-mkvparse"
)

////////////////////////////////////////////////////////////////////////////////

// Handler to extract cover image data
//
// Needs the section parser to handle `AttachmentsElement`
type CoverHandler struct {
	mkvparse.DefaultHandler

	currentAttachmentData     []byte
	currentAttachmentFileName string
	currentAttachmentMIMEType string

	Data     []byte
	MIMEType string
}

func (p *CoverHandler) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	if id == mkvparse.AttachedFileElement && (p.currentAttachmentFileName == "cover.jpg" || p.currentAttachmentFileName == "cover.png") {
		p.Data = p.currentAttachmentData
		p.MIMEType = p.currentAttachmentMIMEType
	}
	return nil
}

func (p *CoverHandler) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	if id == mkvparse.FileNameElement {
		p.currentAttachmentFileName = value
	} else if id == mkvparse.FileMimeTypeElement {
		p.currentAttachmentMIMEType = value
	}
	return nil
}

func (p *CoverHandler) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	if id == mkvparse.FileDataElement {
		p.currentAttachmentData = value
	}
	return nil
}

func main() {}
