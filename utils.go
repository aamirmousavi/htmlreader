package htmlreader

import "golang.org/x/net/html"

func (n *Reader) GetAttribute(key string) (string, bool) {
	for _, attr := range n.HTML.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func (n *Reader) check(attr, val string) bool {
	if n.isElementNode() {
		i, ok := n.GetAttribute(attr)
		if ok && i == val {
			return true
		}
	}
	return false
}

func (n *Reader) isElementNode() bool {
	if n.HTML.Type == html.ElementNode {
		return true
	}
	return false
}
