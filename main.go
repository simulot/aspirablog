package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/simulot/aspirablog/cache"
)

type params struct {
	ConfigFile   string
	BlogName     string
	BlogURL      string
	BlogProvider string
	Destination  string
	TextFormat   string
}

type Application struct {
	config *Config
	cache  *cache.Cache
}

func main() {
	p := params{}
	flag.StringVar(&p.ConfigFile, "config", "config.json", "Configuration file")
	flag.StringVar(&p.BlogURL, "url", "", "Blog's url")
	flag.StringVar(&p.BlogName, "name", "", "Blog's name")
	flag.StringVar(&p.Destination, "destination", "blogs", "Destination folder")
	flag.StringVar(&p.Destination, "provider", "blogger", "Blog provider (blogger)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		check(errors.New("Missing command"))
	}

	c, err := ReadConfigFile(p.ConfigFile)
	if os.IsNotExist(err) {
		c = &Config{
			Folder:       p.Destination,
			BlogURL:      p.BlogURL,
			BlogProvider: p.BlogProvider,
			BlogName:     p.BlogName,
		}
	} else {
		check(err)
	}
	cache, err := cache.New(c.BlogName)
	check(err)
	app := &Application{config: c,
		cache: cache,
	}

	switch flag.Arg(0) {
	case "export":
		err = app.Export(flag.Arg(1))
	}

	check(err)
	err = WriteConfigFile(p.ConfigFile, c)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error as occurred: %s\n", err)
		os.Exit(1)
	}
}
