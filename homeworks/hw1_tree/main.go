package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(output io.Writer, currDir string, printFiles bool) error {
	recursionPrintService("", output, currDir, printFiles)
	return nil
}

func recursionPrintService(prependingString string, output io.Writer, currDir string, printFiles bool) {
	fileObj, err := os.Open(currDir)
	defer fileObj.Close()
	if err != nil {
		log.Fatalf("Could not open %s: %s", currDir, err.Error())
	}
	fileName := fileObj.Name()
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatalf("Could not read dir names in %s: %s", currDir, err.Error())
	}
	var filesMap map[string]os.FileInfo = map[string]os.FileInfo{}
	var unSortedFilesNameArr []string = []string{}
	for _, file := range files {
		unSortedFilesNameArr = append(unSortedFilesNameArr, file.Name())
		filesMap[file.Name()] = file
	}
	sort.Strings(unSortedFilesNameArr)
	var sortedFilesArr []os.FileInfo = []os.FileInfo{}
	for _, stringName := range unSortedFilesNameArr {
		sortedFilesArr = append(sortedFilesArr, filesMap[stringName])
	}
	files = sortedFilesArr
	var newFileList []os.FileInfo = []os.FileInfo{}
	var length int
	if !printFiles {
		for _, file := range files {
			if file.IsDir() {
				newFileList = append(newFileList, file)
			}
		}
		files = newFileList
	}
	length = len(files)
	for i, file := range files {
		if file.IsDir() {
			var stringPrepender string
			if length > i+1 {
				fmt.Fprintf(output, prependingString+"├───"+"%s\n", file.Name())
				stringPrepender = prependingString + "│\t"
			} else {
				fmt.Fprintf(output, prependingString+"└───"+"%s\n", file.Name())
				stringPrepender = prependingString + "\t"
			}
			newDir := filepath.Join(currDir, file.Name())
			recursionPrintService(stringPrepender, output, newDir, printFiles)
		} else if printFiles {
			if file.Size() > 0 {
				if length > i+1 {
					fmt.Fprintf(output, prependingString+"├───%s (%vb)\n", file.Name(), file.Size())
				} else {
					fmt.Fprintf(output, prependingString+"└───%s (%vb)\n", file.Name(), file.Size())
				}
			} else {
				if length > i+1 {
					fmt.Fprintf(output, prependingString+"├───%s (empty)\n", file.Name())
				} else {
					fmt.Fprintf(output, prependingString+"└───%s (empty)\n", file.Name())
				}
			}
		}
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]

	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
