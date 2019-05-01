package mkvparse

import (
	"io"
	"time"
)

// DefaultHandler has basic handlers that do nothing and descend into all
// elements. This can be used as the basis of a handler if you only want to
// implement a subset of the handlers.
type DefaultHandler struct{}

func (dh *DefaultHandler) HandleMasterBegin(ElementID, ElementInfo) (bool, error) {
	return true, nil
}

func (dh *DefaultHandler) HandleMasterEnd(ElementID, ElementInfo) error {
	return nil
}

func (dh *DefaultHandler) HandleString(ElementID, string, ElementInfo) error {
	return nil
}

func (dh *DefaultHandler) HandleInteger(ElementID, int64, ElementInfo) error {
	return nil
}

func (dh *DefaultHandler) HandleFloat(ElementID, float64, ElementInfo) error {
	return nil
}

func (dh *DefaultHandler) HandleDate(ElementID, time.Time, ElementInfo) error {
	return nil
}

func (dh *DefaultHandler) HandleBinary(ElementID, []byte, ElementInfo) error {
	return nil
}

func (dh *DefaultHandler) HandleParseError(reader io.Reader, err *ParseError) *ParseError {
	return err
}
