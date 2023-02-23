package gallery

import (
	"bytes"
	"errors"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

func ThumbnailServer(root http.FileSystem) http.Handler {
	return &thumbHandler{root}
}

type thumbHandler struct {
	root http.FileSystem
}

func (t *thumbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	serveFile(w, r, t.root, path.Clean(upath))
}

func serveFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem, name string) {
	f, err := fs.Open(name)
	if err != nil {
		code := toHTTPError(err)
		http.Error(w, http.StatusText(code), code)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		code := toHTTPError(err)
		http.Error(w, http.StatusText(code), code)
		return
	}

	i, err := thumbnail(f)
	if err != nil {
		code := toHTTPError(err)
		http.Error(w, http.StatusText(code), code)
		return
	}

	var b bytes.Buffer
	if err := jpeg.Encode(&b, i, &jpeg.Options{
		Quality: 10,
	}); err != nil {
		code := toHTTPError(err)
		http.Error(w, http.StatusText(code), code)
		return
	}

	http.ServeContent(w, r, d.Name(), d.ModTime(), bytes.NewReader(b.Bytes()))
}

func toHTTPError(err error) int {
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return http.StatusNotFound
	case errors.Is(err, fs.ErrPermission):
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

func thumbnail(f http.File) (image.Image, error) {
	i, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}

	w, h := float64(i.Bounds().Dx()), float64(i.Bounds().Dy())
	out := image.NewRGBA(image.Rect(0, 0, int(w*(100/h)), 100))
	draw.NearestNeighbor.Scale(out, out.Bounds(), i, i.Bounds(), draw.Over, nil)

	return out, nil
}
