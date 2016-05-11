package main

import (
	"fmt"

	"github.com/kildevaeld/files"
)

func main() {

	fs, _ := files.New("database.bolt")

	file, e := fs.CreatePath("/src/main.go", "./main.go", &files.CreateOptions{
		MkdirP: true,
	})

	if e != nil {
		fmt.Printf("%s\n", e)
		return
	}

	file, e = fs.Get("/src/main.go")

	if e != nil {
		fmt.Printf("%s", e)
	}

	fs.CreateBytes("/test.file", []byte("Hello"), nil)

	fmt.Printf("filename: %s\nSize: %d\nMime: %s\n", file.Filename, file.Size, file.Mime)

	fs.List("/", func(node *files.Node) error {
		fmt.Printf("%#v\n", node)
		return nil
	})

	e = fs.Remove("/src", true)

	fs.ListMeta(func(node string) error {
		fmt.Printf("%#v\n", node)
		return nil
	})
	fmt.Printf("Removed %v\n", e)

}
