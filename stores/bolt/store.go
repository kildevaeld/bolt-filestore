package bolt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/kildevaeld/go-filestore"
	. "github.com/tj/go-debug"
)

var (
	ErrNotExists     = errors.New("ENOENT")
	ErrAlreadyExists = errors.New("EEXIST")
)

var debug = Debug("files:bolt")
var metaBucket = []byte("$meta")
var rootBucket = []byte("/")

type fs_impl struct {
	bolt *bolt.DB
}

func (self *fs_impl) getBucketFromPath(path string, tx *bolt.Tx, create bool, parent bool) (*bolt.Bucket, error) {
	debug("get bucket path: %s", path)

	debug("rootbucket: %s", rootBucket)
	var bucket *bolt.Bucket = tx.Bucket(rootBucket)

	if path == "/" || path == "" {
		return bucket, nil
	}
	if path[0] == '/' {
		path = path[1:]
	}

	split := strings.Split(path, "/")
	l := len(split)
	i := 0

	var err error
	for i < l {

		cur := "/" + split[i]

		debug("	subbucket: %s", cur)
		b := bucket.Bucket([]byte(cur))

		if b == nil {
			debug("	subbucket not exists: %s", cur)
			if !create {
				return nil, ErrNotExists
			}
			debug("	create bucket: %s", cur)
			b, err = bucket.CreateBucket([]byte(cur))
			if err != nil {
				return nil, err
			}
		}
		i++
		if i == l && parent {
			break
		} else {
			bucket = b
		}

	}
	return bucket, nil
}

/*func (self *fs_impl) CreatePath(dest string, src string, options *CreateOptions) (*filestore.File, error) {
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

func (self *fs_impl) CreateBytes(path string, b []byte, options *CreateOptions) (*filestore.File, error) {
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
}*/

func (self *fs_impl) Create(reader io.Reader, file *filestore.File) (uint64, error) {

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return 0, err
	}

	fullPath := filepath.Join(file.Path, file.Filename)

	file.Filesize = uint64(len(bytes))

	err = self.bolt.Update(func(tx *bolt.Tx) error {
		bucket, err := self.getBucketFromPath(file.Path, tx, true, false)
		if err != nil {
			return err
		}

		meta := tx.Bucket(metaBucket)

		b, e := file.Marshal()
		if e != nil {
			return e
		}

		if e := meta.Put([]byte(fullPath), b); e != nil {
			return e
		}

		if e := bucket.Put([]byte(file.Filename), bytes); e != nil {
			meta.Delete([]byte(fullPath))
			return e
		}

		return nil

	})

	if err != nil {
		return 0, err
	}

	return uint64(len(bytes)), nil

	/*if options == nil {
		options = &CreateOptions{}
	}

	if path[0] != '/' {
		path = "/" + path
	}

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	debug("create path: %s, file: %s", dir, filename)

	var file *File

	err := self.bolt.Update(func(tx *bolt.Tx) error {
		bucket, err := self.getBucketFromPath(dir, tx, options.MkdirP, false)
		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		if options.Mime == "" {
			mime, err := DetectContentType(bytes)
			if err != nil || mime == "" {
				mime = "application/octet-stream"
			}
			options.Mime = mime
		}

		if e := bucket.Get([]byte(filename)); e != nil || len(e) > 0 {
			return ErrAlreadyExists
		}

		meta := tx.Bucket(metaBucket)
		now := time.Now().Unix()

		file = &File{
			Fid:      Fid(utils.NewSid()),
			Filename: filename,
			Mime:     options.Mime,
			Filesize: uint64(len(bytes)),
			Path:     dir, //filepath.Join(path, filename),
			Ctime:    now,
			Mtime:    now,
			Perm:     0600,
			Meta:     Meta{},
		}

		b, e := file.Marshal()
		if e != nil {
			return e
		}

		if e := meta.Put([]byte(path), b); e != nil {
			return e
		}

		if e := bucket.Put([]byte(filename), bytes); e != nil {
			meta.Delete([]byte(path))
			return e
		}

		return nil
	})

	return file, err*/
	return 0, nil
}

func (self *fs_impl) Get(path string) (*filestore.File, error) {

	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}

	file := &filestore.File{}
	err := self.bolt.View(func(tx *bolt.Tx) error {

		meta := tx.Bucket(metaBucket)
		if meta == nil {
			return ErrNotExists
		}

		value := meta.Get([]byte(path))

		if value == nil || len(value) == 0 {
			return ErrNotExists
		}

		return file.Unmarshal(value)

	})

	return file, err
}

func (self *fs_impl) Read(path string) (io.Reader, error) {

	var reader *bytes.Buffer

	err := self.bolt.View(func(tx *bolt.Tx) error {
		filename := filepath.Base(path)
		dir := filepath.Dir(path)

		if dir[0] == '/' {
			dir = dir[1:]
		}

		bucket, err := self.getBucketFromPath(dir, tx, false, false)

		if err != nil {
			return err
		}

		b := bucket.Get([]byte(filename))

		if b == nil {
			bucket.ForEach(func(k, v []byte) error {
				fmt.Printf("%s\n", k)
				return nil
			})
			return fmt.Errorf("filename %s not found %s", filename, dir)
		}

		reader = bytes.NewBuffer(b)

		//reader.Seek(0)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return reader, nil
}

func (self *fs_impl) Remove(path string, recursive bool) error {

	_, err := self.Get(path)
	if err != nil {
		if !recursive {
			return err
		}
		return self.bolt.Update(func(tx *bolt.Tx) error {

			bucket, e := self.getBucketFromPath(path, tx, false, true)

			if e != nil {

				return e
			}

			dir := "/" + filepath.Base(path)
			debug("rm bucket: %s", dir)
			e = bucket.DeleteBucket([]byte(dir))

			if e != nil {
				return e
			}

			meta := tx.Bucket(metaBucket)

			cursor := meta.Cursor()
			bPath := []byte(path)
			for k, _ := cursor.Seek(bPath); bytes.HasPrefix(k, bPath); k, _ = cursor.Next() {
				debug("rm meta: %s", k)
				meta.Delete(k)
			}

			return nil

		})

		return err
	}

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	err = self.bolt.Update(func(tx *bolt.Tx) error {

		bucket, err := self.getBucketFromPath(dir, tx, false, false)

		if err != nil {
			return err
		}
		debug("rm file: %s", path)
		err = bucket.Delete([]byte(filename))

		if err != nil {
			return err
		}

		meta := tx.Bucket(metaBucket)
		debug("rm meta: %s", path)
		return meta.Delete([]byte(path))

	})

	return err
}

func buildNodes(prefix string) *filestore.Node {
	root := &filestore.Node{
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
		n := &filestore.Node{
			Dir:    true,
			Parent: root,
			Path:   filepath.Join(root.Path, split[i]),
		}
		i++
		root = n
	}
	return root
}

func (self *fs_impl) List(prefix string, fn func(node *filestore.Node) error) (err error) {

	if prefix == "" || prefix[0] != '/' {
		prefix = "/" + prefix
	}

	parent := buildNodes(prefix)
	debug("parent: %v", parent)
	return self.bolt.View(func(tx *bolt.Tx) error {
		var bucket *bolt.Bucket

		if prefix == "/" {
			bucket = tx.Bucket(rootBucket)
		} else {
			bucket, err = self.getBucketFromPath(prefix, tx, false, false)
			if err != nil {
				return err
			}
		}

		return bucket.ForEach(func(k, v []byte) error {
			node := &filestore.Node{
				Dir:    false,
				Parent: parent,
				Path:   filepath.Join(parent.Path, string(k)),
			}
			if v == nil {
				node.Dir = true
				//parent = node
			}

			return fn(node)
		})

	})

}

func (self *fs_impl) ListMeta(fn func(path string, file filestore.File) error) error {
	return self.bolt.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(metaBucket)
		var file filestore.File
		bucket.ForEach(func(k, v []byte) error {
			file.Unmarshal(v)
			return fn(string(k), file)
		})

		return nil
	})
}

func (self *fs_impl) Mkdir(path string, recursive bool) error {

	/*return self.bolt.Update(func (tx *bolt.Tx) error {

	  })*/

	return nil
}
func (self *fs_impl) Update(path string, info *filestore.File) error {
	file, err := self.Get(path)
	if err != nil {
		return err
	}

	return self.bolt.Update(func(tx *bolt.Tx) error {
		meta := tx.Bucket(metaBucket)

		file.Mtime = time.Now().Unix()
		//file.Perm = mode

		b, e := file.Marshal()
		if e != nil {
			return e
		}

		return meta.Put([]byte(path), b)
	})

	return nil
}

func (self *fs_impl) Close() error {
	return self.bolt.Close()
}

func New(path string) (filestore.FileStore, error) {
	b, e := bolt.Open(path, 0600, nil)
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
