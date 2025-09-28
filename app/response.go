package app

import (
	"strings"

	jquery "github.com/antchfx/jsonquery"
)

type QueryResponse struct {
	Data []string
}

func (q *QueryResponse) AddItem(d string) {
	q.Data = append(q.Data, d)
}

func (q *QueryResponse) String() string {
	return strings.Join(q.Data, "\n")
}

func (r *QueryResponse) Build(n *jquery.Node) {
	if len(n.ChildNodes()) == 1 {
		// This is a leaf node, return its data
		if n.FirstChild != nil {
			r.AddItem(n.FirstChild.Data)
		}
	} else {
		// Recursively process child nodes
		for _, c := range n.ChildNodes() {
			r.Build(c)
		}
	}
}

func (r *QueryResponse) BuildAll(n []*jquery.Node) {
	for _, c := range n {
		r.Build(c)
	}
}

func queryOne(r *jquery.Node, q QueryParams) (resp QueryResponse, err error) {
	n, err := jquery.Query(r, q.Query)
	if err != nil {
		return
	}
	resp.Build(n)
	return
}

func queryAll(r *jquery.Node, q QueryParams) (resp QueryResponse, err error) {
	n, err := jquery.QueryAll(r, q.Query)
	if err != nil {
		return
	}
	resp.BuildAll(n)

	return
}

func runQuery(r *jquery.Node, q QueryParams) (string, error) {
	var resp QueryResponse
	var err error

	if q.Query != "" {
		if q.All {
			resp, err = queryAll(r, q)
		} else {
			resp, err = queryOne(r, q)
		}
		if err != nil {
			return "", err
		}
		return resp.String(), nil
	}
	return r.InnerText(), nil
}
