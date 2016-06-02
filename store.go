package filestore

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/kildevaeld/percy/utils"
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

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

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

			if _, err := seeker.Seek(0, 0); err != nil {
				return nil, err
			}

		} else {
			options.Mime = "application/octet-stream"
		}

	}
	now := time.Now().Unix()
	file := &File{
		Fid:      Fid(utils.NewSid()),
		Filename: filename,
		Mime:     options.Mime,
		//Filesize: uint64(len(bytes)),
		Path:  dir, //filepath.Join(path, filename),
		Ctime: now,
		Mtime: now,
		Perm:  0600,
		Meta:  Meta{},
	}

	size, err := self.store.Create(reader, file)
	if err != nil {
		return nil, err
	}

	file.Filesize = size

	return file, nil
}

func (self *fs_impl) normalizePath(path string) string {
	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}
	return path
}

func (self *fs_impl) Get(path string) (*File, error) {
	path = self.normalizePath(path)
	return self.store.Get(path)
}

func (self *fs_impl) Read(path string) (io.Reader, error) {
	path = self.normalizePath(path)
	return self.store.Read(path)
}

func (self *fs_impl) Remove(path string, recursive bool) error {
	path = self.normalizePath(path)
	return self.store.Remove(path, recursive)
}

func (self *fs_impl) List(prefix string, fn func(node *Node) error) (err error) {
	//path := self.normalizePath(path)
	return self.store.List(prefix, fn)
}

func (self *fs_impl) ListMeta(fn func(path string, file File) error) error {

	return nil
	//return self.store.ListMeta(fn)
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
