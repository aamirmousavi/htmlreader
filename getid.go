package htmlreader

import "golang.org/x/net/html"

func (n *Reader) traverseId(id string) (*html.Node, bool) {
	if n.check("id", id) {
		return n.HTML, true
	}
	for c := n.HTML.FirstChild; c != nil; c = c.NextSibling {
		res, ok := ConvReader(c).traverseId(id)
		if ok {
			return res, true
		}
	}
	return nil, false
}

func (n *Reader) GetElementById(id string) (*html.Node, bool) {
	return n.traverseClass(id)
}
