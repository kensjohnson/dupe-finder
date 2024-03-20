package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listFiles(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Printf("%s - %d bytes\n", path, info.Size())
		return nil
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	dir := os.Args[1]
	err := listFiles(dir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
