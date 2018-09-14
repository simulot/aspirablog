package blogger

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/simulot/aspirablog/http"

	"github.com/simulot/aspirablog/cache"
)

func readTestParam() map[string]string {
	var param = map[string]string{}

	b, err := ioutil.ReadFile("../private/blogger.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &param)
	if err != nil {
		panic(err)
	}
	return param
}

func Test_BlogByURL(t *testing.T) {
	params := readTestParam()
	ca, err := cache.New("test")
	getter := http.New(ca, 10*time.Second)
	service := NewService(getter, params["BloggerKey"])
	if err != nil {
		t.Error(err)
		return
	}

	blog, err := service.ByURL(params["BlogURL"])
	if err != nil {
		t.Error(err)
		return
	}
	if blog.ID != params["BlogID"] {
		t.Errorf("Expecting ID to be %s, got %s", params["BlogID"], blog.ID)
	}
}
