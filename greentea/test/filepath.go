package main

import (
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("D:/", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return err
		}
		println(f.Name())
		return nil
	})

}
