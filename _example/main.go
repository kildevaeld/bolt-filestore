package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kildevaeld/files"
)

func main() {

	fs, _ := files.New("database.bolt")

	r, _ := os.Open("/Users/rasmus/Downloads/qt-creator-opensource-mac-x86_64-4.0.0-beta1.dmg")

	file, err := fs.Create("/min/cool/file.png", r, &files.CreateOptions{
		MkdirP: true,
	})

	fs.CreateBytes("/min/cool/anden-ffile.png", []byte("hello, world"), nil)

	if err != nil {
		fmt.Printf("%v\n", err)
		//return
	} else {
		fmt.Printf("File: %s\nSize: %d\n", file.Filename, file.Size)
	}

	file, err = fs.Get("/min/cool/file.png")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		//return
	} else {
		fmt.Printf("File: %s\nSize: %d\n", file.Filename, file.Size)
	}

	fs.List("/min/cool", func(node *files.Node) error {
		fmt.Printf("%#v\n", node)
		return nil
	})

	o, _ := fs.Read("/min/cool/file.png")

	f, _ := os.Create("test.dmg")
	//fs.Remove("/min/cool/file.png")
	io.Copy(f, o)
}
