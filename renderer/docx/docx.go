package docx

import (
	"log"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
	"github.com/simulot/aspirablog/blog"
)

type HTTPGetter interface {
	IsUpdated(url string) (bool, error)
	Get(url string) ([]byte, string, error)
}

type CacheThere interface {
	IsThere(resource string) (string, error)
	GetResourceFile(resource string) string
}

type HTTPCacher interface {
	HTTPGetter
	CacheThere
}

type Docx struct {
	d   *document.Document
	par document.Paragraph

	run       document.Run
	listLevel int
	http      HTTPCacher
}

func New(http HTTPCacher) *Docx {
	d := &Docx{
		d:    document.New(),
		http: http,
	}
	d.d.AppProperties.SetCompany("responsiveconsulting")
	d.d.AppProperties.SetApplication("Aspirablog")
	d.d.AppProperties.SetApplicationVersion("0.0.0")
	return d
}

func (d *Docx) ConvertBlog(b blog.Blog, file string) error {

	d.Title(b.Title, "Title")

	for _, p := range b.Posts {
		d.AddPost(p)
	}

	return d.d.SaveToFile(file)
}

func (d *Docx) Title(t string, style string) {
	d.AddPar()
	d.par.SetStyle(style)
	d.run = d.par.AddRun()
	d.run.AddText(t)
}

func (d *Docx) AddPost(p blog.Post) {
	d.AddPageBreak()
	d.Title(p.Title, "Heading1")
	d.AddPar()
	d.AddParagraph(p.Paragraph)
}

func (d *Docx) AddParagraph(p blog.Paragraph) {
	for _, f := range p.Fragments {
		switch f.Nature {
		case blog.FragmentText:
			d.AddRun(f.Text, f.Style)
		case blog.FragmentParagraph:
			d.AddParagraph(f.Paragraph)
		case blog.FragmentList:
			d.AddList(f.List)
		case blog.FragmentNewLine:
			d.AddBreak()
		case blog.FragmentImage:
			d.AddImage(f.Image)
		}
	}
}

func (d *Docx) AddPar() {
	d.par = d.d.AddParagraph()
	d.par.Properties().SetAlignment(wml.ST_JcLeft)
}

func (d *Docx) AddRun(t string, s blog.Style) {
	d.run = d.par.AddRun()
	d.run.AddText(t)
	if s.Bold {
		d.run.Properties().SetBold(true)
	}
	if s.Italic {
		d.run.Properties().SetItalic(true)
	}
	return
}

func (d *Docx) AddList(l blog.List) {
	d.listLevel++
	for _, i := range l.Items {
		d.par = d.d.AddParagraph()
		d.par.SetNumberingLevel(d.listLevel - 1)
		d.par.SetNumberingDefinition(d.d.Numbering.Definitions()[0])
		d.AddParagraph(i)
	}
	d.par = d.d.AddParagraph()
	d.listLevel--
}

func (d *Docx) AddBreak() {
	d.AddPar()
	// d.run.AddBreak()
}

func (d *Docx) AddPageBreak() {

}

func (d *Docx) AddImage(image blog.Image) {

	_, err := d.http.IsThere(image.URL)
	if err != nil {
		// Try to load image in cache
		_, _, err := d.http.Get(image.URL)
		if err != nil {
			log.Printf("Can't get image %s: %v", image.URL, err)
			return
		}
	}

	img, err := common.ImageFromFile(d.http.GetResourceFile(image.URL))
	if err != nil {
		log.Printf("unable to create image: %s", err)
		return
	}

	iref, err := d.d.AddImage(img)
	if err != nil {
		log.Printf("unable to add image to document: %s", err)
		return
	}
	d.run = d.par.AddRun()
	inl, err := d.run.AddDrawingInline(iref)
	if err != nil {
		log.Printf("unable to add image to document: %s", err)
	}

	dimension := iref.Size()
	aspectRatio := measurement.Distance(dimension.X) / measurement.Distance(dimension.Y)

	inl.SetSize(15*measurement.Centimeter, 15*measurement.Centimeter/aspectRatio)

}
