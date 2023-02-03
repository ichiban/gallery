package gallery

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

type Image struct {
	URL string `json:"url"`
}

func List(images *[]Image, root string) error {
	*images = []Image{}

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		path, err = filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		switch ext := strings.ToLower(filepath.Ext(path)); ext {
		case ".jpg", ".jpeg", ".gif", ".png":
			log.Printf("image: %s", path)
			*images = append(*images, Image{
				URL: filepath.Join("/files", path),
			})
		default:
			log.Printf("skip: %s", path)
		}

		return nil
	})
}
