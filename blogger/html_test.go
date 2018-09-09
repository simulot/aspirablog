package blogger

import (
	"testing"

	"golang.org/x/net/html"
)

func Test_nodeStyle(t *testing.T) {
	type args struct {
		n     *html.Node
		style string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"center",
			args{
				n: &html.Node{
					Attr: []html.Attribute{
						{Namespace: "", Key: "style", Val: "color: blue; align: center; background:"},
					},
				},
				style: "align",
			},
			"center",
		},
		{
			"color",
			args{
				n: &html.Node{
					Attr: []html.Attribute{
						{Namespace: "", Key: "style", Val: "color: blue; align: center; background:"},
					},
				},
				style: "color",
			},
			"blue",
		},
		{
			"color",
			args{
				n: &html.Node{
					Attr: []html.Attribute{
						{Namespace: "", Key: "style", Val: "color: blue; align: center; background:"},
					},
				},
				style: "background",
			},
			"",
		},
		{
			"forecolor",
			args{
				n: &html.Node{
					Attr: []html.Attribute{
						{Namespace: "", Key: "style", Val: "color: blue; align: center; background:"},
					},
				},+
				+96
				style: "forecolor",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodeStyle(tt.args.n, tt.args.style); got != tt.want {
				t.Errorf("nodeStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}
