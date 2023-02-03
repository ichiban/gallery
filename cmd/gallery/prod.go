//go:build release
// +build release

package main

import (
	"io/fs"
	"log"
	"net/http"
)

func root() http.Handler {
	pub, err := fs.Sub(gallery.UIBuild, "ui/build")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(pub))
}
