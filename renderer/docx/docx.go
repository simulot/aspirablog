package docx

import (
	"io/ioutil"
	"log"
	"time"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
	"github.com/simulot/aspirablog/blog"
	qrcode "github.com/skip2/go-qrcode"
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
	d.d.Settings.SetUpdateFieldsOnOpen(true) // Force TOC to update on opening
	d.CoverPage(b)

	list := b.Posts //[len(b.Posts)-10:]
	for i, p := range list {
		d.AddPost(p)
		if i < len(list)-1 {
			d.AddPageBreak()
		}
	}
	log.Println("Writing document to file", file)
	return d.d.SaveToFile(file)
}

func (d *Docx) CoverPage(b blog.Blog) {
	d.AddPar()
	d.Title(b.Title, "Title")

	d.Title("Date de mise à jour :", "Subtitle")
	d.AddRun(time.Now().Format("02/01/2006"), blog.Style{})

	d.Title("Adresse du blog", "Subtitle")
	d.AddRun(b.URL, blog.Style{})
	d.AddBreak()
	d.addQRCode(b.URL)

	d.Title("Menu", "Subtitle")
	d.AddPar()
	// Add a TOC
	d.d.AddParagraph().AddRun().AddField(document.FieldTOC)
	d.AddPageBreak()

}

func (d *Docx) Title(t string, style string) {
	d.AddPar()
	d.par.SetStyle(style)
	d.run = d.par.AddRun()
	d.run.AddText(t)
}

func (d *Docx) AddPost(p blog.Post) {
	log.Printf("Writing post '%s'\n", p.Title)
	d.Title(p.Published.Format("02/01/06")+", "+p.Title, "Heading1")
	d.AddPar()
	if len(p.Categories) > 0 {
		d.AddPar()
		d.AddRun("Catégories : ", blog.Style{Bold: true})
		for i, c := range p.Categories {
			if i == len(p.Categories)-1 {
				d.AddRun(c+".", blog.Style{})
			} else {
				d.AddRun(c+", ", blog.Style{})
			}
		}
	}
	d.AddPar()
	d.AddParagraph(p.Paragraph)
	d.addURLandQRCode(p.URL)

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
	switch s.Align {
	case blog.AlignLeft:
		d.par.Properties().SetAlignment(wml.ST_JcLeft)
	case blog.AlignCenter:
		d.par.Properties().SetAlignment(wml.ST_JcCenter)
	case blog.AlignRight:
		d.par.Properties().SetAlignment(wml.ST_JcRight)
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
	d.run.AddBreak()
}

func (d *Docx) AddPageBreak() {
	d.d.AddParagraph().Properties().AddSection(wml.ST_SectionMarkNextPage)
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

	err = d.addImageFile(d.http.GetResourceFile(image.URL), true)
	if err != nil {
		log.Printf("unable to create image: %s", err)
	}

}

func (d *Docx) addImageFile(filename string, adjustSize bool) error {
	img, err := common.ImageFromFile(filename)
	if err != nil {
		return err
	}

	iref, err := d.d.AddImage(img)
	if err != nil {
		return err
	}
	d.run = d.par.AddRun()
	inl, err := d.run.AddDrawingInline(iref)
	if err != nil {
		return err
	}

	if adjustSize {
		dimension := iref.Size()
		aspectRatio := measurement.Distance(dimension.X) / measurement.Distance(dimension.Y)
		inl.SetSize(15*measurement.Centimeter, 15*measurement.Centimeter/aspectRatio)
	}
	return nil
}

func (d *Docx) addURLandQRCode(url string) {
	d.AddPar()
	d.par.Properties().SetAlignment(wml.ST_JcCenter)
	d.addQRCode(url)
	d.AddPar()
	d.par.Properties().SetAlignment(wml.ST_JcCenter)
	d.run.AddText(url)
	d.par.Properties().SetAlignment(wml.ST_JcLeft)
	d.AddPar()
}

func (d *Docx) addQRCode(text string) {

	qrFile, err := ioutil.TempFile("", "aspirablog*.png")
	if err != nil {
		log.Fatalf("Can't create temporay file: %v", err)
	}
	qrName := qrFile.Name()
	png, err := qrcode.Encode(text, qrcode.Medium, 128)
	if err != nil {
		qrFile.Close()
		log.Fatalf("Can't create qrcode: %v", err)
	}
	_, err = qrFile.Write(png)
	if err != nil {
		qrFile.Close()
		log.Fatalf("Can't create qrcode: %v", err)
	}
	qrFile.Close()
	err = d.addImageFile(qrName, false)
	if err != nil {
		log.Fatalf("Can't create qrcode: %v", err)
	}
}
