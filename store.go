package filestore

import (
	"bytes"
	"io"
	"os"

	. "github.com/tj/go-debug"
)

var debug = Debug("files")

var metaBucket = []byte("$meta")
var rootBucket = []byte("/")

type fs_impl struct {
	store FileStore
}

func (self *fs_impl) CreatePath(dest string, src string, options *CreateOptions) (*File, error) {
	if options == nil {
		options = &CreateOptions{}
	}

	if options.Mime == "" {
		mime, err := DetectContentTypeFromPath(src)
		if err != nil || mime == "" {
			mime = "application/octet-stream"
		}
		options.Mime = mime
	}

	reader, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return self.Create(dest, reader, options)
}

func (self *fs_impl) CreateBytes(path string, b []byte, options *CreateOptions) (*File, error) {
	buf := bytes.NewBuffer(b)

	if options == nil {
		options = &CreateOptions{}
	}

	if options.Mime == "" {
		mime, err := DetectContentType(b)
		if err != nil || mime == "" {
			mime = "application/octet-stream"
		}
		options.Mime = mime
	}

	return self.Create(path, buf, options)
}

func (self *fs_impl) Create(path string, reader io.Reader, options *CreateOptions) (*File, error) {
	if options == nil {
		options = &CreateOptions{}
	}

	if path[0] != '/' {
		path = "/" + path
	}

	//filename := filepath.Base(path)
	//dir := filepath.Dir(path)

	if options.Mime == "" {

		if seeker, ok := reader.(io.ReadSeeker); ok {
			var bytes [256]byte
			if _, err := seeker.Read(bytes[:]); err != nil {
				return nil, err
			}

			mime, err := DetectContentType(bytes[:])
			if err != nil || mime == "" {
				mime = "application/octet-stream"
			}
			options.Mime = mime

		} else {
			options.Mime = "application/octet-stream"
		}

	}

	return nil, nil
}

func (self *fs_impl) Get(path string) (*File, error) {

	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}

	return self.store.Get(path)

}

func (self *fs_impl) Read(path string) (io.Reader, error) {
	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}

	return self.store.Read(path)
}

func (self *fs_impl) Remove(path string, recursive bool) error {
	return nil
}

func (self *fs_impl) List(prefix string, fn func(node *Node) error) (err error) {
	return nil
}

func (self *fs_impl) ListMeta(fn func(path string, file File) error) error {
	return nil
}

func (self *fs_impl) Mkdir(path string, recursive bool) error {

	/*return self.bolt.Update(func (tx *bolt.Tx) error {

	  })*/

	return nil
}

func (self *fs_impl) Chmod(path string, mode FileMode) error {

	/*file, err := self.Get(path)
	if err != nil {
		return err
	}*/

	return nil

}
func (self *fs_impl) Chown(path string, uid []byte) error {
	/*file, err := self.Get(path)
	if err != nil {
		return err
	}*/

	return nil

}
func (self *fs_impl) Chgrp(path string, guid []byte) error {
	return nil
}

func (self *fs_impl) SetMeta(path string, metadata Meta) error {
	return nil
}

func New(store FileStore) (FS, error) {
	return &fs_impl{store}, nil
}
