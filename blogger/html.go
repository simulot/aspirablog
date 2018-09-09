package blogger

import (
	"log"
	"regexp"
	"strings"

	"github.com/simulot/aspirablog/blog"
	"golang.org/x/net/html"
)

func (b *Blogger) parsePost(s string) (blog.Paragraph, error) {
	p := blog.Paragraph{}

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return p, err
	}
	style := blog.Style{}
	p = b.parseParagraph(doc, style)

	return p, nil
}

func (b *Blogger) parseParagraph(n *html.Node, style blog.Style) blog.Paragraph {
	p := blog.NewParagraph()

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			p.AddText(c.Data, style)
			continue
		}
		if c.Type == html.ElementNode {
			switch c.Data {
			case "div":
				p.AddParagraph(b.parseParagraph(c, style))
			case "b":
				s := style
				s.Bold = true
				p.AddParagraph(b.parseParagraph(c, s))
			case "i":
				s := style
				s.Italic = true
				p.AddParagraph(b.parseParagraph(c, s))
			case "br":
				p.AddNewLine()
			case "a":
				p.AddParagraph(b.parseParagraph(c, style))
			case "img":
				p.AddImage(b.parseImage(c))
			case "ol", "ul":
				p.AddList(b.parseList(c, style))
			case "body", "html", "table", "tbody", "td":
				p.AddParagraph(b.parseParagraph(c, style))
			case "tr":
				p.AddParagraph(b.parseParagraph(c, style))
				p.AddNewLine()
				// case "iframe":
				// 	p.AddParagraph(b.parseIframe(c, sytle))
			}
		}
	}
	return p
}

func (b *Blogger) parseImage(n *html.Node) blog.Image {
	image := blog.Image{}

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			image.URL = attr.Val
		}
	}

	parent := n.Parent
	if parent != nil && parent.Type == html.ElementNode && parent.Data == "a" {
		for _, attr := range parent.Attr {
			if attr.Key == "href" {
				image.URL = attr.Val
			}
		}
	}
	if image.URL == "" {
		return image
	}

	_, _, err := b.http.Get(image.URL)
	if err != nil {
		log.Printf("URL %s cached: %v", image.URL, err)
	}
	return image
}

func (b *Blogger) parseList(n *html.Node, style blog.Style) blog.List {
	list := blog.NewList(n.Data == "ol")
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "li" {
			list.AddItem(b.parseParagraph(c, style))
		}
	}
	return list
}

func (b *Blogger) parseIframe(n *html.Node, style blog.Style) blog.Paragraph {
	p := blog.Paragraph{}
	if b, class := nodeAttribut(n, "class"); b {
		switch class {
		case "YOUTUBE-iframe-video":

		}
	}
	return p
}

func nodeAttribut(n *html.Node, attr string) (bool, string) {
	for _, t := range n.Attr {
		if t.Key == attr {
			return true, t.Val
		}
	}
	return false, ""
}

var reStyle = regexp.MustCompile(`([[:alnum:]_-]+):\s*([^;]*)`)

func nodeStyle(n *html.Node, style string) string {
	if b, c := nodeAttribut(n, "style"); b {
		mm := reStyle.FindAllStringSubmatch(c, -1)
		for _, m := range mm {
			if len(m) == 3 && m[1] == style {
				return m[2]
			}
		}
	}
	return ""
}
