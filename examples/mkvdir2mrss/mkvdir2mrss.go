package main

////////////////////////////////////////////////////////////////////////////////
// This example parses all MKV files in a dir, and generates a (Media) RSS feed
// for them. Cover art is extracted, and put in the same dir as the output feed.
//
// Usage:
//     ./mkvdir2mrss --baseURL http://localhost --out=feeds/feed.xml Movies/

////////////////////////////////////////////////////////////////////////////////

import (
	"crypto/sha1"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/remko/go-mkvparse"
)

////////////////////////////////////////////////////////////////////////////////
// Parsing
////////////////////////////////////////////////////////////////////////////////

type MediaFile struct {
	title    string
	artist   string
	duration time.Duration
	cover    []byte
}

type MediaParser struct {
	mkvparse.DefaultHandler

	duration                  float64
	timecodeScale             int64
	currentTagGlobal          bool
	currentTagName            *string
	currentTagValue           *string
	currentAttachmentData     []byte
	currentAttachmentFileName string
	mediaFile                 *MediaFile
}

func (p *MediaParser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	if id == mkvparse.TagElement {
		p.currentTagGlobal = true
	} else if id == mkvparse.SimpleTagElement {
		p.currentTagName = nil
		p.currentTagValue = nil
	}
	return true, nil
}

func (p *MediaParser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	if id == mkvparse.SimpleTagElement && p.currentTagGlobal && p.currentTagName != nil && p.currentTagValue != nil {
		if *p.currentTagName == mkvparse.TagArtist {
			p.mediaFile.artist = *p.currentTagValue
		}
	} else if id == mkvparse.AttachedFileElement {
		if p.currentAttachmentFileName == "cover.jpg" {
			p.mediaFile.cover = p.currentAttachmentData
		}
	}
	return nil
}

func (p *MediaParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	if id == mkvparse.TagNameElement {
		p.currentTagName = &value
	} else if id == mkvparse.TagStringElement {
		p.currentTagValue = &value
	} else if id == mkvparse.TitleElement {
		p.mediaFile.title = value
	} else if id == mkvparse.FileNameElement {
		p.currentAttachmentFileName = value
	}
	return nil
}

func (p *MediaParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	if (id == mkvparse.TagTrackUIDElement || id == mkvparse.TagEditionUIDElement || id == mkvparse.TagChapterUIDElement || id == mkvparse.TagAttachmentUIDElement) && value != 0 {
		p.currentTagGlobal = false
	} else if id == mkvparse.TimecodeScaleElement {
		p.timecodeScale = value
	}
	return nil
}

func (p *MediaParser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	if id == mkvparse.DurationElement {
		p.duration = value
	}
	return nil
}

func (p *MediaParser) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	if id == mkvparse.FileDataElement {
		p.currentAttachmentData = value
	}
	return nil
}

var supportedMediaFileRE = regexp.MustCompile(`(?i)\.mk[av]$`)
var unsupportedMediaFileRE = regexp.MustCompile(`(?i)\.(mp4|m4v|avi|mpg)$`)

func parseFile(path string) (*MediaFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	handler := MediaParser{
		duration:      -1.0,
		timecodeScale: 1000000,
		mediaFile: &MediaFile{
			duration: -1,
		},
	}
	err = mkvparse.ParseSections(file, &handler, mkvparse.InfoElement, mkvparse.TagsElement, mkvparse.TracksElement, mkvparse.AttachmentsElement)
	if err != nil {
		return nil, err
	}

	if handler.duration >= 0 {
		handler.mediaFile.duration = time.Duration(int64(handler.duration * float64(handler.timecodeScale)))
	} else {
		handler.mediaFile.duration = -1
	}
	return handler.mediaFile, nil
}

////////////////////////////////////////////////////////////////////////////////
// RSS Generation
////////////////////////////////////////////////////////////////////////////////

type RSSMediaContent struct {
	XMLName  xml.Name `xml:"media:content"`
	URL      string   `xml:"url,attr"`
	FileSize int64    `xml:"fileSize,attr"`
	Duration int      `xml:"duration,attr"`
	Type     string   `xml:"type,attr"`
}

type RSSEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  int64    `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
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

func run() error {
	var noDirs bool
	baseURL := flag.String("baseURL", "", "Base URL")
	outFile := flag.String("out", "", "Output RSS Feed file")
	flag.BoolVar(&noDirs, "noDirs", false, "Don't include dirnames in titles")
	flag.Parse()
	dirs := flag.Args()

	if baseURL == nil || outFile == nil || len(dirs) < 1 {
		return fmt.Errorf("Missing parameters")
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
				return fmt.Errorf("Error walking %s: %v", path, err)
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
				if err != nil {
					return fmt.Errorf("Error loading %s: %v", path, err)
				} else {
					item := &RSSItem{
						Title:          name,
						PubDate:        info.ModTime().Format(time.RFC822),
						ITunesDuration: formatDuration(file.duration),
						Author:         file.artist,
						MediaContent: &RSSMediaContent{
							FileSize: info.Size(),
							Duration: int(file.duration / time.Second),
							URL:      mediaURL,
							Type:     "video/x-matroska",
						},
					}
					item.Enclosure = &RSSEnclosure{
						URL:    item.MediaContent.URL,
						Length: item.MediaContent.FileSize,
						Type:   item.MediaContent.Type,
					}
					if len(file.artist) > 0 {
						item.MediaCredit = &RSSMediaCredit{
							Role:  "author",
							Value: file.artist,
						}
					}
					if len(file.cover) > 0 {
						thumbFile := filepath.Join(outDir, fmt.Sprintf("%x.jpg", sha1.Sum(file.cover)))
						if _, err := os.Stat(thumbFile); os.IsNotExist(err) {
							err := ioutil.WriteFile(thumbFile, file.cover, 0644)
							if err != nil {
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

	output, err := xml.MarshalIndent(feed, "  ", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(*outFile, output, 0644)
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
