package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	startDepth := strings.Count(dir, "\\") + 2

	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			last := strings.Split(path, "\\")


			if len(last) == startDepth - 1{
				fmt.Println(last[len(last)-1])
			} else {
				spaces := strings.Repeat(" ", (len(last) - startDepth) * 6)
				fmt.Println(spaces, "├────", last[len(last)-1])
			}
			return nil
		})

	if err != nil {
		log.Println(err)
	}
}
