package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ichiban/gallery"
)

func main() {
	var (
		port int
		dir  string
	)
	flag.IntVar(&port, "p", 8080, "http listen port")
	flag.StringVar(&dir, "d", ".", "root directory")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/images", func(w http.ResponseWriter, r *http.Request) {
		var images []gallery.Image
		if err := gallery.List(&images, dir); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(images)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(b)
	})
	r.Mount("/files/", http.StripPrefix("/files/", http.FileServer(http.FS(os.DirFS(dir)))))
	r.Mount("/", root())

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	log.Print(err)
}
