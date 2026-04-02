package query

import (
	jquery "github.com/antchfx/jsonquery"
	"github.com/bnassif/jquery-proxy/pkg/server/params"
	"github.com/bnassif/jquery-proxy/pkg/server/response"
)

type Query struct {
	Data   []byte
	Params *params.Params
}

func (q *Query) queryOne(r *jquery.Node) (resp response.Response, err error) {
	n, err := jquery.Query(r, q.Params.Query)
	if err != nil {
		return
	}
	resp.Build(n)
	return
}

func (q *Query) queryAll(r *jquery.Node) (resp response.Response, err error) {
	n, err := jquery.QueryAll(r, q.Params.Query)
	if err != nil {
		return
	}
	for _, c := range n {
		resp.Build(c)
	}

	return
}

func (q *Query) Run() (string, error) {

	var resp response.Response

	if q.Params.Query == "" {
		return string(q.Data), nil
	}

	// If a query was provided, create a jQuery node for parsing
	r, err := parseToJqueryNode(q.Data)
	if err != nil {
		return "", nil
	}

	// Run jQuery Parsing
	if q.Params.All {
		resp, err = q.queryAll(r)
	} else {
		resp, err = q.queryOne(r)
	}
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
