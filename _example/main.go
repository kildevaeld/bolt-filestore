package main

import (
	"fmt"
	"os"

	"github.com/kildevaeld/files"
)

func print(indent string, node *files.Node, fs files.FS) error {
	if node.Dir {
		fmt.Printf("%sDirectory: %s\n", indent, node.Path)
		return fs.List(node.Path, func(node *files.Node) error {

			return print(indent+"  ", node, fs)
		})
	}

	file, e := fs.Get(node.Path)
	if e != nil {
		return e
	}
	fmt.Printf("%sFile: %s, Size: %d, Mime: %s, Id: %s, Perm: %s\n", indent, file.Filename, file.Filesize, file.Mime, file.Fid.Hex(), file.Perm)
	return nil
}

func main() {

	fs, _ := files.New("database.bolt")
	defer os.Remove("database.bolt")
	fs.CreatePath("/src/main.go", "./main.go", &files.CreateOptions{
		MkdirP: true,
	})

	fs.CreateBytes("/src/test.txt", []byte("Hello, You wonderful woman"), nil)

	/*file, e = fs.Get("/src/main.go")

	if e != nil {
		fmt.Printf("%s", e)
	}*/

	fs.CreateBytes("/test.txt", []byte("Hello"), nil)
	fs.CreateBytes("/test2.txt", []byte("Hello, world"), nil)
	fs.CreateBytes("/src/lib/another.html", []byte("<html></html>"), &files.CreateOptions{
		MkdirP: true,
	})

	fs.List("/", func(node *files.Node) error {
		return print("", node, fs)
	})
	//fmt.Printf("LIST %v", e)
	e := fs.Remove("/src/lib", true)
	fmt.Println("\nRoot: /")
	fs.List("/", func(node *files.Node) error {
		return print(" ", node, fs)
	})

	fmt.Printf("Removed %v\n", e)

}
