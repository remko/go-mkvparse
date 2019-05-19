////////////////////////////////////////////////////////////
// Extracts the cover from a .mkv file, and displays it
////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"github.com/remko/go-mkvparse"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"time"
)

type handler struct {
	currentAttachmentData     []byte
	currentAttachmentFileName string
	currentAttachmentMIMEType string
	cover                     []byte
	coverMIMEType             string
}

func (p *handler) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	return true, nil
}

func (p *handler) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	if id == mkvparse.AttachedFileElement && p.currentAttachmentFileName == "cover.jpg" {
		p.cover = p.currentAttachmentData
		p.coverMIMEType = p.currentAttachmentMIMEType
	}
	return nil
}

func (p *handler) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	if id == mkvparse.FileNameElement {
		p.currentAttachmentFileName = value
	} else if id == mkvparse.FileMimeTypeElement {
		p.currentAttachmentMIMEType = value
	}
	return nil
}

func (p *handler) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	return nil
}

func (p *handler) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	return nil
}

func (p *handler) HandleDate(id mkvparse.ElementID, value time.Time, info mkvparse.ElementInfo) error {
	return nil
}

func (p *handler) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	if id == mkvparse.FileDataElement {
		p.currentAttachmentData = value
	}
	return nil
}

func run() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()
	h := handler{}
	if err := mkvparse.ParseSections(file, []mkvparse.ElementID{mkvparse.AttachmentsElement}, &h); err != nil {
		return err
	}
	if len(h.cover) == 0 {
		return fmt.Errorf("no cover found")
	}

	// Decode image
	var img image.Image
	if h.coverMIMEType == "image/jpeg" {
		img, err = jpeg.Decode(bytes.NewReader(h.cover))
	} else if h.coverMIMEType == "image/png" {
		img, err = png.Decode(bytes.NewReader(h.cover))
	} else {
		return fmt.Errorf("unknown MIME type: %s", h.coverMIMEType)
	}
	if err != nil {
		return err
	}

	// Display image
	dw, err := wde.NewWindow(img.Bounds().Max.X-img.Bounds().Min.X, img.Bounds().Max.Y-img.Bounds().Min.Y)
	if err != nil {
		return err
	}
	dw.SetTitle("Cover")
	draw.Draw(dw.Screen(), img.Bounds(), img, image.ZP, draw.Src)
	dw.FlushImage()
	dw.Show()
	go func() {
		events := dw.EventChan()
		for {
			select {
			case event := <-events:
				switch event.(type) {
				case wde.CloseEvent:
					wde.Stop()
				}
			}
		}
	}()
	wde.Run()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
