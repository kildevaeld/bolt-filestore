package files

import (
	"errors"
	"io"
)

var (
	ErrNotExists     = errors.New("ENOENT")
	ErrAlreadyExists = errors.New("EEXIST")
)

/*type File struct {
	Filename string
	Size     uint64
	Mime     string
	Ctime    time.Time
	c
}*/

type CreateOptions struct {
	Overwrite bool
	MkdirP    bool
	Mime      string
	Perm      FileMode
}

type Node struct {
	Parent *Node
	Dir    bool
	Path   string
}

func (self *Node) Root() *Node {
	node := self
	for {
		if node.Parent == nil {
			break
		}
		node = node.Parent
	}
	return node
}

type FS interface {
	CreatePath(dest string, src string, options *CreateOptions) (*File, error)
	CreateBytes(path string, b []byte, options *CreateOptions) (*File, error)
	Create(path string, reader io.Reader, options *CreateOptions) (*File, error)
	Read(path string) (io.Reader, error)
	Get(path string) (*File, error)
	Remove(path string, recursive bool) error
	//Mkdir(path string, recursive bool) error
	Chmod(path string, mode FileMode) error
	Chown(path string, uid []byte) error
	Chgrp(path string, guid []byte) error
	List(prefix string, fn func(node *Node) error) error
	ListMeta(fn func(path string, file File) error) error
}
