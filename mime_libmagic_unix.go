// +build libmagic

package filestore

import "github.com/rakyll/magicmime"

func detectContentType(sample []byte) (string, error) {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		return "", err
	}
	defer magicmime.Close()

	return magicmime.TypeByBuffer(sample)

}

func detectContentTypeFromPath(path string) (string, error) {

	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		return "", err
	}

	mimetype, err := magicmime.TypeByFile(path)
	if err != nil {
		return "", err
	}

	magicmime.Close()

	return mimetype, nil
}
