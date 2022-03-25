package htmlreader

import (
	"strings"

	"golang.org/x/net/html"
)

type (
	Reader struct {
		HTML *html.Node
	}
)

func New(str string) (*Reader, error) {
	r := strings.NewReader(str)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return &Reader{
		HTML: doc,
	}, nil
}

func convReader(n *html.Node) *Reader {
	return &Reader{
		HTML: n,
	}
}
