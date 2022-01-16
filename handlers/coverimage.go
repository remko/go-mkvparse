package mkvparse

import (
	"bytes"
	"fmt"
	"image"
	png "image/jpeg"
	jpeg "image/png"
)

func ParseCoverImage(path string) (image.Image, error) {
	data, typ, err := ParseCover(path)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	if typ == "image/jpeg" {
		return jpeg.Decode(bytes.NewReader(data))
	} else if typ == "image/png" {
		return png.Decode(bytes.NewReader(data))
	} else {
		return nil, fmt.Errorf("unknown MIME type: %s", typ)
	}
}
