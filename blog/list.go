package blog

type List struct {
	Ordered bool
	Level   int
	Items   []Paragraph
}

func NewList(ordered bool) List {
	return List{
		Ordered: ordered,
		Items:   []Paragraph{},
	}
}

func (l *List) AddItem(p Paragraph) {
	l.Items = append(l.Items, p)
}
