package gallery

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

type Image struct {
	URL      string `json:"url"`
	ThumbURL string `json:"thumb_url"`
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
		case ".jpg", ".jpeg":
			log.Printf("image: %s", path)
			*images = append(*images, Image{
				URL:      filepath.Join("/files", path),
				ThumbURL: filepath.Join("/thumbs", path),
			})
		default:
			log.Printf("skip: %s", path)
		}

		return nil
	})
}
