package htmlreader

import (
	"bytes"
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

func (n *Reader) ToString() (string, error) {
	var b bytes.Buffer
	err := html.Reader(n.HTML)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
