// mkvdir2mrss parses all MKV files in a dir, and generates a (Media) RSS feed
// for them. Cover art is extracted, and put in the same dir as the output feed.
//
// Usage:
//     ./mkvdir2mrss --baseURL http://localhost --out=feeds/feed.xml Movies/
package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/remko/go-mkvparse"
	"golang.org/x/image/draw"
)

////////////////////////////////////////////////////////////////////////////////
// Parsing
////////////////////////////////////////////////////////////////////////////////

type MediaFile struct {
	Title             string
	Artist            string
	Duration          time.Duration
	Cover             []byte
	Channels          int64
	HasVideo          bool
	HasAudio          bool
	SamplingFrequency float64
}

type MediaParser struct {
	mkvparse.DefaultHandler

	duration          float64
	timecodeScale     int64
	title             string
	channels          int64
	hasVideo          bool
	hasAudio          bool
	samplingFrequency float64
}

func (p *MediaParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TitleElement:
		p.title = value
	}
	return nil
}

func (p *MediaParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TimecodeScaleElement:
		p.timecodeScale = value
	case mkvparse.ChannelsElement:
		if value > p.channels {
			p.channels = value
		}
	case mkvparse.TrackTypeElement:
		switch value {
		case mkvparse.TrackTypeVideo:
			p.hasVideo = true
		case mkvparse.TrackTypeAudio:
			p.hasAudio = true
		}
	}
	return nil
}

func (p *MediaParser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.DurationElement:
		p.duration = value
	case mkvparse.SamplingFrequencyElement:
		if value > p.samplingFrequency {
			p.samplingFrequency = value
		}
	}
	return nil
}

func parseFile(path string) (*MediaFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	handler := MediaParser{
		duration:      -1.0,
		channels:      0,
		timecodeScale: 1000000,
	}
	tagsh := mkvparse.NewTagsHandler()
	coverh := mkvparse.CoverHandler{}
	err = mkvparse.ParseSections(file, mkvparse.NewHandlerChain(tagsh, &coverh, &handler), mkvparse.InfoElement, mkvparse.TagsElement, mkvparse.TracksElement, mkvparse.AttachmentsElement)
	if err != nil {
		return nil, err
	}

	mf := MediaFile{
		Title:             handler.title,
		Artist:            tagsh.Tags()[mkvparse.TagArtist],
		Cover:             coverh.Data,
		Channels:          handler.channels,
		HasAudio:          handler.hasAudio,
		HasVideo:          handler.hasVideo,
		SamplingFrequency: handler.samplingFrequency,
	}
	if handler.duration >= 0 {
		mf.Duration = time.Duration(int64(handler.duration * float64(handler.timecodeScale)))
	} else {
		mf.Duration = -1
	}
	return &mf, nil
}

////////////////////////////////////////////////////////////////////////////////
// RSS Generation
////////////////////////////////////////////////////////////////////////////////

type RSSMediaContent struct {
	XMLName      xml.Name `xml:"media:content"`
	URL          string   `xml:"url,attr"`
	FileSize     int64    `xml:"fileSize,attr"`
	Duration     int      `xml:"duration,attr"`
	Channels     int64    `xml:"channels,attr,omitempty"`
	Type         string   `xml:"type,attr,omitempty"`
	Medium       string   `xml:"medium,attr,omitempty"`
	SamplingRate string   `xml:"samplingrate,attr,omitempty"`
}

type RSSEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  int64    `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type RSSMediaTitle struct {
	XMLName xml.Name `xml:"media:title"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type RSSMediaCredit struct {
	XMLName xml.Name `xml:"media:credit"`
	Role    string   `xml:"role,attr"`
	Value   string   `xml:",chardata"`
}

type RSSMediaThumbnail struct {
	XMLName xml.Name `xml:"media:thumbnail"`
	URL     string   `xml:"url,attr"`
}

type RSSItem struct {
	XMLName   xml.Name `xml:"item"`
	PubDate   string   `xml:"pubDate"`
	Title     string   `xml:"title"`
	Author    string   `xml:"author,omitempty"`
	Enclosure *RSSEnclosure

	MediaContent   *RSSMediaContent
	MediaCredit    *RSSMediaCredit
	MediaTitle     *RSSMediaTitle
	MediaThumbnail *RSSMediaThumbnail

	ITunesDuration string `xml:"itunes:duration"`
}

type RSSChannel struct {
	XMLName xml.Name   `xml:"channel"`
	Title   string     `xml:"title"`
	Items   []*RSSItem `xml:"items"`
}

type RSSFeed struct {
	XMLName  xml.Name    `xml:"rss"`
	Version  string      `xml:"version,attr"`
	MediaNS  string      `xml:"xmlns:media,attr"`
	ITunesNS string      `xml:"xmlns:itunes,attr"`
	Channel  *RSSChannel `xml:"channel"`
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

var supportedMediaFileRE = regexp.MustCompile(`(?i)\.mk[av]$`)
var unsupportedMediaFileRE = regexp.MustCompile(`(?i)\.(mp4|m4v|avi|mpg)$`)

func run() error {
	var noDirs bool
	baseURL := flag.String("baseURL", "", "Base URL")
	outFile := flag.String("out", "", "Output RSS Feed file")
	flag.BoolVar(&noDirs, "noDirs", false, "Don't include dirnames in titles")
	flag.Parse()
	dirs := flag.Args()

	if baseURL == nil || outFile == nil || len(dirs) < 1 {
		return fmt.Errorf("missing parameters")
	}

	baseDir, _ := os.Getwd()
	outDir, outFilename := filepath.Split(*outFile)
	outDir, err := filepath.Abs(outDir)
	if err != nil {
		return err
	}
	extension := filepath.Ext(outFilename)
	title := outFilename[0 : len(outFilename)-len(extension)]

	feed := RSSFeed{
		Version:  "2.0",
		MediaNS:  "http://search.yahoo.com/mrss/",
		ITunesNS: "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Channel: &RSSChannel{
			Title: title,
		},
	}

	for _, dir := range dirs {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			return err
		}
		err = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("error walking %s: %v", path, err)
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			publicPath, err := filepath.Rel(baseDir, path)
			if err != nil {
				return err
			}
			filename, err := filepath.Rel(absDir, path)
			if err != nil {
				return err
			}
			extension := filepath.Ext(filename)
			name := filename[0 : len(filename)-len(extension)]
			if noDirs {
				name = filepath.Base(name)
			}
			mediaURL := fmt.Sprintf("%s/%s", *baseURL, strings.Replace(url.PathEscape(publicPath), "%2F", "/", -1))

			if supportedMediaFileRE.MatchString(path) {
				file, err := parseFile(path)
				if len(file.Title) == 0 {
					file.Title = name
				}
				if err != nil {
					return fmt.Errorf("error loading %s: %v", path, err)
				} else {
					item := &RSSItem{
						Title:          file.Title,
						PubDate:        info.ModTime().Format(time.RFC822),
						ITunesDuration: formatDuration(file.Duration),
						Author:         file.Artist,
						MediaContent: &RSSMediaContent{
							FileSize: info.Size(),
							Duration: int(file.Duration / time.Second),
							URL:      mediaURL,
						},
						MediaTitle: &RSSMediaTitle{
							Type:  "plain",
							Value: file.Title,
						},
					}
					item.Enclosure = &RSSEnclosure{
						URL:    item.MediaContent.URL,
						Length: item.MediaContent.FileSize,
						Type:   item.MediaContent.Type,
					}
					if len(file.Artist) > 0 {
						item.MediaCredit = &RSSMediaCredit{
							Role:  "author",
							Value: file.Artist,
						}
					}
					if len(file.Cover) > 0 {
						thumbFile := filepath.Join(outDir, fmt.Sprintf("%x.jpg", sha1.Sum(file.Cover)))
						if _, err := os.Stat(thumbFile); os.IsNotExist(err) {
							img, err := scale(file.Cover, 512)
							if err != nil {
								return err
							}
							if err := ioutil.WriteFile(thumbFile, img, 0644); err != nil {
								return err
							}
						}
						publicThumbFile, err := filepath.Rel(baseDir, thumbFile)
						if err != nil {
							return err
						}
						item.MediaThumbnail = &RSSMediaThumbnail{
							URL: fmt.Sprintf("%s/%s", *baseURL, strings.Replace(url.PathEscape(publicThumbFile), "%2F", "/", -1)),
						}
					}
					if file.Channels > 0 {
						item.MediaContent.Channels = file.Channels
					}
					if file.HasVideo {
						item.MediaContent.Medium = "video"
						item.MediaContent.Type = "video/x-matroska"
					} else if file.HasAudio {
						item.MediaContent.Medium = "audio"
						item.MediaContent.Type = "audio/x-matroska"
					}
					if file.SamplingFrequency > 0 {
						item.MediaContent.SamplingRate = strconv.FormatFloat(float64(file.SamplingFrequency)/1000, 'f', -1, 64)
					}
					feed.Channel.Items = append(feed.Channel.Items, item)
				}
			} else if unsupportedMediaFileRE.MatchString(path) {
				// Fallback to basic information for unsupported media files
				item := &RSSItem{
					Title:   name,
					PubDate: info.ModTime().Format(time.RFC822),
					Enclosure: &RSSEnclosure{
						URL:    mediaURL,
						Length: info.Size(),
						Type:   "video/mp4",
					},
				}
				feed.Channel.Items = append(feed.Channel.Items, item)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	output, err := xml.Marshal(feed)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(*outFile, output, 0644)
}

func scale(data []byte, size int) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	width := size
	height := size
	aspect := float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())
	if aspect > float64(width)/float64(height) {
		height = int(float64(height) / aspect)
	} else {
		width = int(float64(width) * aspect)
	}
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	out := bytes.Buffer{}
	err = jpeg.Encode(&out, dst, &jpeg.Options{Quality: 75})
	return out.Bytes(), err
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
