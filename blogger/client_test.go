package blogger

import (
	"testing"
	"time"

	"github.com/simulot/aspirablog/cache"
	"github.com/simulot/aspirablog/http"
)

func Test_GetAll(t *testing.T) {
	params := readTestParam()
	ca, err := cache.New("test")
	getter := http.New(ca, 10*time.Second)

	client := New("test", params["BlogURL"], params["BloggerKey"], getter)
	l, err := client.GetAll()

	_, _ = l, err
}
