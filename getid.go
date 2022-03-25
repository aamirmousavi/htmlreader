package htmlreader

import "golang.org/x/net/html"

func (n *Reader) GetElementById(id string) (*html.Node, bool) {
	return n.traverse("id", id)
}
