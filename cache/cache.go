package cache

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Cache struct {
	path string
}

func New(path string) (*Cache, error) {
	userTmp, err := os.UserCacheDir()
	if err != nil {
		return nil, errors.Wrapf(err, "Can't create cache directory")
	}
	path = filepath.Join(userTmp, "aspirablog", path)
	err = os.MkdirAll(path, 0700)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't create cache directory")
	}

	return &Cache{
		path: path,
	}, nil

}

func (c *Cache) IsThere(resource string) (string, error) {
	n := c.GetResourceFile(resource) + ".etag"

	b, err := ioutil.ReadFile(n)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Cache) Store(resource string, b []byte, etag string) error {
	n := c.GetResourceFile(resource)
	err := ioutil.WriteFile(n, b, 0600)
	if err != nil {
		return errors.Wrapf(err, "Can't store in cache")
	}
	err = ioutil.WriteFile(n+".etag", []byte(etag), 0600)
	if err != nil {
		return errors.Wrapf(err, "Can't store in cache")
	}
	return nil
}

func (c *Cache) Load(resource string) ([]byte, string, error) {
	n := c.GetResourceFile(resource)

	etag, err := ioutil.ReadFile(n + ".etag")
	if err != nil {
		return nil, "", err
	}

	b, err := ioutil.ReadFile(n)
	if err != nil {
		return nil, "", err
	}
	return b, string(etag), nil
}

func (c *Cache) GetResourceFile(resource string) string {
	return filepath.Join(c.path, c.hashResourceName(resource))
}

func (c *Cache) hashResourceName(n string) string {
	hash := sha1.New()
	hash.Write([]byte(n))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
