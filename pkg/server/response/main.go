package response

import (
	"strings"

	jquery "github.com/antchfx/jsonquery"
)

type Response struct {
	Content []string
}

func (r *Response) addItem(d string) {
	r.Content = append(r.Content, d)
}

func (r *Response) String() string {
	return strings.Join(r.Content, "\n")
}

func (r *Response) Build(n *jquery.Node) {
	if len(n.ChildNodes()) == 1 {
		// This is a leaf node, return its data
		if n.FirstChild != nil {
			r.addItem(n.FirstChild.Data)
		}
	} else {
		// Recursively process child nodes
		for _, c := range n.ChildNodes() {
			r.Build(c)
		}
	}
}
