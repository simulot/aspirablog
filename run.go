package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/simulot/aspirablog/blog"
	"github.com/simulot/aspirablog/blogger"
	"github.com/simulot/aspirablog/http"
	"github.com/simulot/aspirablog/renderer/docx"
)

type BlogProvider interface {
	GetAll() (blog.Blog, error)
}
type Renderer interface {
	ConvertBlog(b blog.Blog, file string) error
}

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

func (a *Application) Export(format string) error {
	httpClient := http.New(a.cache, 10*time.Second)
	err := os.MkdirAll(filepath.Join(a.config.Folder, a.config.BlogName), 0700)
	check(errors.Wrapf(err, "Application can't Export"))
	p, err := a.NewProvider(httpClient)

	check(errors.Wrapf(err, "Application can't Export"))
	b, err := a.Load(p)
	check(errors.Wrapf(err, "Application can't Export"))

	r, err := a.NewRenderer(format, httpClient)
	check(errors.Wrapf(err, "Application can't Export"))
	f := filepath.Join(a.config.Folder, a.config.BlogName, b.Title+".docx")
	r.ConvertBlog(b, f)

	return nil
}

func (a *Application) Load(p BlogProvider) (blog.Blog, error) {
	return p.GetAll()
}

func (a *Application) NewProvider(http HTTPGetter) (BlogProvider, error) {

	switch a.config.BlogProvider {
	case "blogger":
		log.Printf("Reading blog at %s\n", a.config.BlogURL)
		return blogger.New(a.config.BlogName, a.config.BlogURL, a.config.BloggerAPIKey, http), nil
	}
	return nil, errors.Errorf("Unsupported blog provider: %s", a.config.BlogProvider)
}

func (a *Application) NewRenderer(format string, http HTTPCacher) (Renderer, error) {
	switch format {
	case "docx":
		return docx.New(http), nil
	}
	return nil, errors.Errorf("Renderer: unsupported format:%s", format)
}
