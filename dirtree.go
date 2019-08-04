package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
// TODO: file extensions info, hyperlink (only newer consoles)
// TODO: colored levels?
// TODO: file info?
// TODO: handling non directory files
// TODO: color by file extensions?
// TODO: all system, i guess easy and provided by go
// TODO: ???
func printFileEntry(file os.FileInfo, indent uint32, indentMask uint32, separator string) {
	var sb strings.Builder
	for i := indent; i > 0; i-- {
		shift := i - 1
		if indentMask & (1 << shift) != 0 {
			sb.WriteString("│    ")
		} else {
			sb.WriteString("     ")
		}
	}
	fmt.Println(sb.String() + separator + file.Name())
}

func printDirTree(directory string, indent uint32, indentMask uint32) error{
	iterator, err := ioutil.ReadDir(directory)

	for i, file := range iterator {
		if file.IsDir() {
			var newIndentMask uint32
			if i == len(iterator) - 1 {
				printFileEntry(file, indent, indentMask,"└────")
				newIndentMask = indentMask << 1
			} else {
				printFileEntry(file, indent, indentMask,"├────")
				newIndentMask = (indentMask << 1) | 1
			}
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
		fmt.Println(err.Error())
	}

	if err := printDirTree(dir, 0, 0); err != nil {
		fmt.Println(err)
	}
}