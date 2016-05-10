package files

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/boltdb/bolt"
)

var metaBucket = []byte("$meta")
var rootBucket = []byte("/")

type fs_impl struct {
	bolt *bolt.DB
}

func (self *fs_impl) getBucketFromPath(path string, tx *bolt.Tx, create bool) (*bolt.Bucket, error) {

	split := strings.Split(path, "/")
	l := len(split)
	i := 0

	bucket := tx.Bucket(rootBucket)

	if path == "/" || path == "" {
		return bucket, nil
	}

	var err error
	for i < l {
		cur := split[i]
		if cur == "" {
			break
		}

		b := bucket.Bucket([]byte(cur))
		if b == nil {
			if !create {
				return nil, ErrNotExists
			}
			b, err = bucket.CreateBucket([]byte(cur))
			if err != nil {
				return nil, err
			}
		}

		bucket = b

		i++
	}
	return bucket, nil
}

func (self *fs_impl) CreateBytes(path string, b []byte, options *CreateOptions) (*File, error) {
	buf := bytes.NewBuffer(b)
	return self.Create(path, buf, options)
}

func (self *fs_impl) Create(path string, reader io.Reader, options *CreateOptions) (*File, error) {
	if options == nil {
		options = &CreateOptions{}
	}

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	if dir[0] == '/' {
		dir = dir[1:]
	}

	var file *File

	err := self.bolt.Update(func(tx *bolt.Tx) error {
		bucket, err := self.getBucketFromPath(dir, tx, options.MkdirP)
		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		if e := bucket.Get([]byte(filename)); e != nil || len(e) > 0 {
			return errors.New("Already exists")
		}

		if e := bucket.Put([]byte(filename), bytes); e != nil {
			return e
		}

		meta := tx.Bucket(metaBucket)

		file = &File{
			Filename: filename,
			Mime:     "application/octet-stream",
			Size:     uint64(len(bytes)),
		}

		b, _ := json.Marshal(file)

		return meta.Put([]byte(path), b)
	})

	return file, err
}

func (self *fs_impl) Get(path string) (*File, error) {

	var file File
	err := self.bolt.View(func(tx *bolt.Tx) error {

		meta := tx.Bucket(metaBucket)
		if meta == nil {
			return ErrNotExists
		}

		value := meta.Get([]byte(path))

		if value == nil || len(value) == 0 {
			return ErrNotExists
		}

		return json.Unmarshal(value, &file)
	})

	return &file, err
}

func (self *fs_impl) Read(path string) (io.Reader, error) {

	var reader io.Reader

	err := self.bolt.View(func(tx *bolt.Tx) error {
		filename := filepath.Base(path)
		dir := filepath.Dir(path)

		if dir[0] == '/' {
			dir = dir[1:]
		}

		bucket, err := self.getBucketFromPath(dir, tx, false)

		if err != nil {
			return err
		}

		b := bucket.Get([]byte(filename))

		reader = bytes.NewBuffer(b)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return reader, nil
}

func (self *fs_impl) Remove(path string) error {

	_, err := self.Get(path)
	if err != nil {
		return err
	}

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	err = self.bolt.Update(func(tx *bolt.Tx) error {

		bucket, err := self.getBucketFromPath(dir, tx, false)

		if err != nil {
			return err
		}

		err = bucket.Delete([]byte(filename))

		if err != nil {
			return err
		}

		meta := tx.Bucket(metaBucket)

		return meta.Delete([]byte(path))

	})

	return err
}

func buildNodes(prefix string) *Node {
	root := &Node{
		Dir:    true,
		Parent: nil,
		Path:   "/",
	}

	if prefix == "" || prefix == "/" {
		return root
	}

	split := strings.Split(prefix, "/")
	l := len(split)
	i := 0

	for i < l {
		n := &Node{
			Dir:    true,
			Parent: root,
			Path:   filepath.Join(root.Path, split[i]),
		}
		i++
		root = n
	}
	return root
}

func (self *fs_impl) List(prefix string, fn func(node *Node) error) error {

	if prefix[0] == '/' {
		prefix = prefix[1:]
	}

	parent := buildNodes(prefix)

	self.bolt.Update(func(tx *bolt.Tx) error {

		bucket, err := self.getBucketFromPath(prefix, tx, false)
		if err != nil {
			return err
		}

		bucket.ForEach(func(k, v []byte) error {
			node := &Node{
				Dir:    false,
				Parent: parent,
				Path:   filepath.Join(parent.Path, string(k)),
			}
			if v == nil {
				node.Dir = true
				parent = node
			}

			return fn(node)
		})

		return nil
	})

	return nil
}

func (self *fs_impl) Mkdir(path string, recursive bool) error {

	/*return self.bolt.Update(func (tx *bolt.Tx) error {

	  })*/

	return nil
}

func New(path string) (FS, error) {
	b, e := bolt.Open(path, 0766, nil)
	if e != nil {
		return nil, e
	}

	err := b.Update(func(tx *bolt.Tx) error {

		if _, e := tx.CreateBucketIfNotExists(metaBucket); e != nil {
			return e
		}

		if _, e := tx.CreateBucketIfNotExists(rootBucket); e != nil {
			return e
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &fs_impl{b}, nil
}
