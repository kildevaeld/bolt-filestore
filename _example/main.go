package main

import (
	"encoding/json"
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

	fmt.Printf("%sFile: %s, Size: %d, Mime: %s, Id: %s, Perm: %s -- %v\n", indent, file.Filename, file.Filesize, file.Mime, file.Fid.Hex(), file.Perm, file.Meta)
	return nil
}

func main() {

	fs, _ := filestore.New("database.bolt")
	defer os.Remove("database.bolt")
	fs.CreatePath("/src/main.go", "./main.go", &filestore.CreateOptions{
		MkdirP: true,
	})

	fs.SetMeta("/src/main.go", filestore.Meta{
		"Hello": "World",
	})

	fs.Chmod("/src/main.go", 0704)

	fs.CreateBytes("/src/test.txt", []byte("Hello, You wonderful woman"), nil)

	/*file, e = fs.Get("/src/main.go")

	if e != nil {
		fmt.Printf("%s", e)
	}*/

	fs.CreateBytes("/test.txt", []byte("Hello"), nil)
	fs.CreateBytes("/test2.txt", []byte("Hello, world"), nil)
	fs.CreateBytes("/src/lib/another.html", []byte("<html></html>"), &filestore.CreateOptions{
		MkdirP: true,
	})

	fs.List("/", func(node *filestore.Node) error {
		return print("", node, fs)
	})
	//fmt.Printf("LIST %v", e)
	e := fs.Remove("/src/lib", true)
	fmt.Println("\nRoot: /")
	fs.List("/", func(node *filestore.Node) error {
		return print(" ", node, fs)
	})

	fmt.Printf("Removed %v\n", e)

	file, _ := fs.Get("/src/test.txt")

	b, _ := json.Marshal(file)
	fmt.Printf("%s", b)
}
