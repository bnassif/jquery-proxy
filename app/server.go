package app

import (
	"fmt"
	"net/http"
)

func makeServer(c ServerConfig) *http.Server {

	return &http.Server{
		Addr:         c.AddressPort,
		ReadTimeout:  getDuration(c.ReadTimeout),
		WriteTimeout: getDuration(c.WriteTimeout),
		IdleTimeout:  getDuration(c.IdleTimeout),
		// TODO: Add a custom handler
		Handler: http.HandlerFunc(handleRequest),
	}
}

func internalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	q := QueryParams{}

	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("Request URI: %s\n", r.URL.String())

	// Parse the received request
	err := q.parseFromURL(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Parsed Query Parameters: %+v\n", q)

	// Make the request
	resBody, err := request(q.URL.String())
	if err != nil {
		internalError(w, err)
		return
	}

	// Parse the response
	resp, err := parseResponse(resBody)
	if err != nil {
		internalError(w, err)
		return
	}

	// Query the response, if specified and return a string
	s, err := runQuery(resp, q)
	if err != nil {
		internalError(w, err)
		return
	}

	// Write the response
	status, err := w.Write([]byte(s))
	if err != nil {
		internalError(w, err)
		return
	}

	fmt.Printf("Response Status: %d\n", status)
}
