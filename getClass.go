package htmlreader

import "golang.org/x/net/html"

func (n *Reader) traverseClass(class string) (*html.Node, bool) {
	if n.check("class", class) {
		return n.HTML, true
	}
	for c := n.HTML.FirstChild; c != nil; c = c.NextSibling {
		res, ok := convReader(c).traverseClass(class)
		if ok {
			return res, true
		}
	}
	return nil, false
}

func (n *Reader) GetElmentByClass(class string) (*html.Node, bool) {
	return n.traverseClass(class)
}
