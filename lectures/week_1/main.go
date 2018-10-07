package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prew string
	for in.Scan() {
		txt := in.Text()

		if txt == prew {
			continue
		}

		if txt < prew {
			return fmt.Errorf("file not sorted")
		}

		prew = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := unig(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
