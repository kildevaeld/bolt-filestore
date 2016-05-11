// +build windows
package filestore

import "bitbucket.org/taruti/mimemagic"

func detectContentType(sample []byte) (string, error) {
	return mimemagic.Match("application/octet-stream", sample), nil
}

func detectContentTypeFromPath(path string) (string, error) {
	return "application/octet-stream", nil
}
