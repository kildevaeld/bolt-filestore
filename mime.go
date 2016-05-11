package filestore

import (
	"mime"
	"path/filepath"
)

func DetectContentType(sample []byte) (string, error) {
	return detectContentType(sample)
}

func DetectContentTypeFromPath(path string) (string, error) {
	ext := filepath.Ext(path)
	m := mime.TypeByExtension(ext)
	if m == "" {
		return detectContentTypeFromPath(path)
	}
	return m, nil
}
