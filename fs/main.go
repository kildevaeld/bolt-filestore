package main

import (
	"fmt"
	"os"

	"github.com/kildevaeld/go-filestore"
)

func print(indent string, node *filestore.Node, fs filestore.FS) error {
	if node.Dir {
		fmt.Printf("%sDirectory: %s\n", indent, node.Path)
		return fs.List(node.Path, func(node *filestore.Node) error {

			return print(indent+"  ", node, fs)
		})
	}

	file, e := fs.Get(node.Path)
	if e != nil {
		return e
	}

	fmt.Printf("%sFile: %s, Size: %d, Mime: %s, Id: %s, Perm: %s -- %v\n", indent, file.Filename, file.Filesize, file.Mime, file.Fid.Hex(), file.Perm, file.Path)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <dbpath>", os.Args[0])
		return
	}

	path := os.Args[1]

	fs, e := filestore.New(path)

	if e != nil {
		fmt.Printf("%v\n", e)
	}

	fs.List("/", func(node *filestore.Node) error {
		return print("", node, fs)
	})

}
