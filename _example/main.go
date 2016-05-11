package main

import (
	"fmt"

	"github.com/kildevaeld/files"
)

func main() {

	fs, _ := files.New("database.bolt")

	fs.CreatePath("/src/main.go", "./main.go", &files.CreateOptions{
		MkdirP: true,
	})

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
		fmt.Printf("%#v\n", node)
		return nil
	})

	e := fs.Remove("/src/lib", true)

	fs.ListMeta(func(node string) error {
		fmt.Printf("%#v\n", node)
		return nil
	})

	fs.List("/src", func(node *files.Node) error {
		fmt.Printf("%#v\n", node)
		return nil
	})

	fmt.Printf("Removed %v\n", e)

}
