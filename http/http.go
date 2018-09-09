package http

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/simulot/aspirablog/cache"
)

type Client struct {
	*cache.Cache
	http.Client
}

func New(cache *cache.Cache, httpTimeout time.Duration) *Client {
	return &Client{
		Cache: cache,
		Client: http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   httpTimeout,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
			},
			Timeout: httpTimeout,
		},
	}
}

// IsUpdated queries the server end return True when the server has
// newer version than  provided ETAG
func (c *Client) IsUpdated(url string) (bool, error) {
	etag, err := c.Cache.IsThere(url)
	if err != nil {
		return true, nil
	}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return false, errors.Wrap(err, "Can't check HEAD")
	}
	req.Header.Add("If-None-Match", etag)

	resp, err := c.Do(req)

	if err != nil {
		return false, errors.Wrap(err, "Can't check HEAD")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode == http.StatusNotModified {
		return false, nil
	}
	if resp.StatusCode != http.StatusOK {
		return false, errors.Errorf("Can't check HEAD: got %s", resp.Status)
	}
	return true, nil
}

// Get performs a GET to the given url. It returns a buffer, the last etag and error
func (c *Client) Get(url string) ([]byte, string, error) {
	if updated, err := c.IsUpdated(url); !updated && err == nil {
		b, etag, err := c.Cache.Load(url)
		if err == nil {
			return b, etag, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", errors.Wrapf(err, "Can't GET %s", url)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, "", errors.Wrapf(err, "Can't GET %s", url)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, resp.Header.Get("ETag"), errors.Wrapf(err, "Can't GET %s", url)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		c.Cache.Store(url, b, resp.Header.Get("ETag"))
	}

	return b, resp.Header.Get("ETag"), errors.Wrapf(err, "Can't GET %s", url)

}
