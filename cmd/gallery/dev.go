//go:build dev
// +build dev

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func root() http.Handler {
	u, err := url.Parse("http://localhost:5173/")
	if err != nil {
		log.Fatal(err)
	}
	return httputil.NewSingleHostReverseProxy(u)
}
