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

	fmt.Printf("filename: %s\nSize: %d\nMime: %s\n", file.Filename, file.Size, file.Mime)

}
