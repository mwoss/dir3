package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func printFileEntry(file os.FileInfo, indent uint32, indentMask uint32, separator string) {
	var sb strings.Builder
	for i := uint32(0); i <= indent; i++ {
		if indentMask & (1 << i) != 0{
			sb.WriteString("│    ")
		} else{
			sb.WriteString("     ")
		}
	}
	fmt.Println(sb.String(), separator, file.Name())
}

func printDirTree(directory string, indent uint32, indentMask uint32) error{
	iterator, err := ioutil.ReadDir(directory)

	for i, file := range iterator {
		if file.IsDir() {
			printFileEntry(file, indent, indentMask,"├────")
			var newIndentMask uint32

			if i == len(iterator) - 1 {
				newIndentMask = (indentMask << 1) | 1
			} else {
				newIndentMask = indentMask << 1
			}
			fmt.Println("INDENT MASK: ", newIndentMask)
			err = printDirTree(filepath.Join(directory, file.Name()), indent + 1, newIndentMask)
		} else if i == len(iterator) - 1{
			printFileEntry(file, indent, indentMask, "└────")
		} else {
			printFileEntry(file, indent, indentMask, "├────")
		}
	}
	return err
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if err := printDirTree(dir, 0, 0); err != nil {
		log.Fatal(err)
	}
}
