package query

import (
	"bytes"
	"fmt"

	jquery "github.com/antchfx/jsonquery"
	"github.com/bnassif/jquery-proxy/pkg/server/params"
)

func parseToJqueryNode(b []byte) (res *jquery.Node, err error) {
	res, err = jquery.Parse(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("Error parsing response: %s\n", err)
	}
	return
}

func NewQuery(d []byte, p *params.Params) Query {
	return Query{
		Data:   d,
		Params: p,
	}
}
