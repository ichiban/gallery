//go:build !dev
// +build !dev

package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/ichiban/gallery"
)

func root() http.Handler {
	pub, err := fs.Sub(gallery.UIBuild, "ui/build")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(pub))
}
