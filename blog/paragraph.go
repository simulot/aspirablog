package blog

func NewParagraph() Paragraph {
	return Paragraph{
		Fragments: []Fragment{},
	}
}
func (p *Paragraph) AddText(t string, s Style) {
	if len(t) > 0 {
		f := Fragment{
			Nature: FragmentText,
			Text:   TextLinter(t),
			Style:  s,
		}
		p.Fragments = append(p.Fragments, f)
	}
}

func (p *Paragraph) AddNewLine() {
	if len(p.Fragments) > 1 && p.Fragments[len(p.Fragments)-2].Nature == FragmentNewLine && p.Fragments[len(p.Fragments)-1].Nature == FragmentNewLine {
		return
	}
	f := Fragment{
		Nature: FragmentNewLine,
	}
	p.Fragments = append(p.Fragments, f)
}
func (p *Paragraph) AddParagraph(paragraph Paragraph) {
	f := Fragment{
		Nature:    FragmentParagraph,
		Paragraph: paragraph,
	}
	p.Fragments = append(p.Fragments, f)
}

func (p *Paragraph) AddImage(i Image) {
	if i.URL != "" {
		p.Fragments = append(p.Fragments, Fragment{
			Nature: FragmentImage,
			Image:  i,
		})
	}
}

func (p *Paragraph) AddList(l List) {
	if len(l.Items) > 0 {
		p.Fragments = append(p.Fragments, Fragment{
			Nature: FragmentList,
			Style:  p.Style,
			List:   l,
		})
	}
}

func (p *Paragraph) AddVideo(v Video) {
	if v.Thumbnail != "" {
		p.Fragments = append(p.Fragments, Fragment{
			Nature: FragmentVideo,
			Style:  p.Style,
			Video:  v,
		})
	}
}
