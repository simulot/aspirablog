package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	BlogName      string
	BlogURL       string
	BlogProvider  string
	LastETAG      string
	Folder        string
	ImageFormat   string
	TextFormat    string
	BloggerAPIKey string
}

func ReadConfigFile(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrap(err, "ReadConfigFile")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "ReadConfigFile")
	}

	c := &Config{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, errors.Wrap(err, "ReadConfigFile")
	}

	return c, nil
}

func WriteConfigFile(file string, c *Config) error {
	f, err := os.Create(file)

	if err != nil {
		return errors.Wrap(err, "WriteConfigFile")
	}

	defer f.Close()

	b, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "WriteConfigFile")
	}

	_, err = f.Write(b)

	if err != nil {
		return errors.Wrap(err, "WriteConfigFile")
	}
	return nil
}
