package htmlreader

import "golang.org/x/net/html"

func (n *Reader) GetElmentByClass(class string) (*html.Node, bool) {
	return n.traverse("class", class)
}
