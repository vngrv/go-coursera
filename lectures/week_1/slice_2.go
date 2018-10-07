package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	// get slice
	sl1 := buf[1:4] // [2,3,4]
	sl2 := buf[:2]  // [1,2]
	sl3 := buf[2:]  // [3,4,5]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:] // [1,2,3,4,5]
	newBuf[0] = 9

	// buf = [9 2 3 4 5 6]
	// newBuf = [1 2 3 4 5 6]
	newBuf = append(newBuf, 6)
	fmt.Println(newBuf)

	// copy slice in other slice
	var emptyBuf []int // len = 0, cap = 0
	// wrong - will copy the smaller of the two
	copied := copy(emptyBuf, buf)
	fmt.Println(emptyBuf, copied)

	// right
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) // ints = [1,5,6,4]
	fmt.Println(ints)
}
