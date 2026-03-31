package params

import (
	"fmt"
	"net/url"
)

type Params struct {
	URL   url.URL
	Query string
	All   bool
}

func NewParams(u *url.URL) (params *Params, err error) {
	var reqUrl *url.URL
	var reqQuery string
	var reqAll bool

	p := u.Query()

	// Ensure the url parameter is present
	if !p.Has("url") {
		err = fmt.Errorf("missing required parameter: url")
		return
	}

	// Unescape the URL value to a raw string
	rawUrl, err := url.QueryUnescape(p.Get("url"))
	if err != nil {
		return
	}

	// Parse the raw URL string
	reqUrl, err = url.Parse(rawUrl)
	if err != nil {
		return
	}

	// Parse the optional `query` and `contains` parameters
	if p.Has("query") {
		reqQuery, err = url.QueryUnescape(p.Get("query"))
		if err != nil {
			return nil, err
		}

		// Assess the `all` parameter if query is present
		reqAll = p.Has("all") && p.Get("all") == "true"
	}

	return &Params{
		URL:   *reqUrl,
		Query: reqQuery,
		All:   reqAll,
	}, nil
}
