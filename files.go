package files

import (
	"errors"
	"io"
)

var (
	ErrNotExists = errors.New("ENOENT")
)

type File struct {
	Filename string
	Size     uint64
	Mime     string
}

type CreateOptions struct {
	Overwrite bool
	MkdirP    bool
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
	CreateBytes(path string, b []byte, options *CreateOptions) (*File, error)
	Create(path string, reader io.Reader, options *CreateOptions) (*File, error)
	Read(path string) (io.Reader, error)
	Get(path string) (*File, error)
	Remove(path string) error
	Mkdir(path string, recursive bool) error
	List(prefix string, fn func(node *Node) error) error
}
