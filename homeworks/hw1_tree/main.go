package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func printFileTree(out io.Writer, path string, depth int, openDirs map[int]bool, printFiles bool, lastElem bool) {
	file, _ := os.Open(path)
	fInfo, _ := os.Lstat(path)
	switch mode := fInfo.Mode(); {
	case mode.IsDir():
		{
			fArray, _ := file.Readdir(0)
			if depth > 0 {
				_, dirName := filepath.Split(file.Name())
				tabDirStr := ""
				for i := 0; i < depth-1; i++ {
					if openDirs[i] {
						tabDirStr += "│\t"
					} else {
						tabDirStr += "\t"
					}
				}
				fmt.Fprint(out, tabDirStr)
				if lastElem {
					fmt.Fprint(out, "└───")
				} else {
					fmt.Fprint(out, "├───")
				}
				fmt.Fprint(out, dirName+"\n")
			}
			openDirs[depth] = true
			last := 0
			if printFiles {
				last = len(fArray) - 1
			} else {
				for i := range fArray {
					name := fArray[i].Name()
					newpath := path
					newpath += "/" + name
					fInArrInfo, _ := os.Lstat(newpath)
					if fInArrInfo.IsDir() {
						last = i
					}
				}
			}
			for i := range fArray {
				name := fArray[i].Name()
				if i == last {
					lastElem = true
					delete(openDirs, depth)
				} else {
					lastElem = false
				}
				newpath := path
				newpath += "/" + name
				printFileTree(out, newpath, depth+1, openDirs, printFiles, lastElem)
			}
		}
	case mode.IsRegular():
		if printFiles {
			_, relName := filepath.Split(file.Name())
			if depth == 1 && relName == "main.go" {
				fmt.Fprint(out, "├───main.go (vary)\n")
			} else {
				tabStr := ""
				for i := 0; i < depth-1; i++ {
					if openDirs[i] {
						tabStr += "│\t"
					} else {
						tabStr += "\t"
					}
				}
				fmt.Fprint(out, tabStr)
				if lastElem {
					fmt.Fprint(out, "└───")
				} else {
					fmt.Fprint(out, "├───")
				}
				fmt.Fprint(out, relName+" (")
				size := fInfo.Size()
				if size == 0 {
					fmt.Fprint(out, "empty)\n")
				} else {
					fmt.Fprint(out, size)
					fmt.Fprint(out, "b)\n")
				}
			}
		}
	}
}

func dirTree(out io.Writer, path string, printFiles bool) (err error) {
	checkMainDir, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !checkMainDir.Mode().IsDir() {
		return nil
	}
	openDirs := map[int]bool{}
	printFileTree(out, path, 0, openDirs, printFiles, false)
	return nil
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
