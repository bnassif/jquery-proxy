package app

import (
	"fmt"
	"net/url"
)

type QueryParams struct {
	URL   url.URL
	Query string
	All   bool
}

func (q *QueryParams) parseFromURL(u *url.URL) (err error) {
	// Parse the URL Query Parameters
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
	u, err = url.Parse(rawUrl)
	if err != nil {
		return
	}

	// Set the URL
	q.URL = *u

	// Parse the optional `query` and `contains` parameters
	if p.Has("query") {
		q.Query, err = url.QueryUnescape(p.Get("query"))
		if err != nil {
			return
		}

		// Assess the `all` parameter if query is present
		if p.Has("all") && p.Get("all") == "true" {
			q.All = true
		}
	}

	return
}
