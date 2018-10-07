package main

import "fmt"

func main() {
	// create
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)

	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)

	// someInt := buf2[0]

	// add elements
	var buf []int
	buf = append(buf, 9, 10)
	buf = append(buf, 12)

	otherBuf := make([]int, 3) // [0,0,0]
	buf = append(buf, otherBuf...)

	//
	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)

}
