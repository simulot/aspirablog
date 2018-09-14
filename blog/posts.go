package blog

import (
	"time"
)

// Post structure
type Post struct {
	Title      string    // Post title
	Author     string    // Post author
	Categories []string  // Post categories
	Published  time.Time // Post publication time
	URL        string    // HTML url
	Paragraph  Paragraph // Content
}

// Paragraph structure
type Paragraph struct {
	Style     Style
	Fragments []Fragment
}

// Style structure hold text style
type Style struct {
	Align     Alignment
	Bold      bool
	Italic    bool
	Underline bool
}

// Alignment type enumerates all kind of alignment
type Alignment int

// Alignment type enumerates all kind of alignment
const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

type FragmentNature int

const (
	FragmentText FragmentNature = iota
	FragmentList
	FragmentImage
	FragmentParagraph
	FragmentNewLine
	FragmentVideo
)

type Fragment struct {
	Nature FragmentNature
	Style  Style

	Text      string
	Image     Image
	List      List
	Paragraph Paragraph
	Video     Video
}
type Image struct {
	URL   string
	Title string
}

type Video struct {
	URL       string
	Thumbnail string
}

func New() *Post {
	return &Post{}
}

func (p *Post) SetTitle(s string) *Post {
	p.Title = s
	return p
}

func (p *Post) SetPubished(t time.Time) *Post {
	p.Published = t
	return p
}
