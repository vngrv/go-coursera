package main

import "fmt"

func main() {
	var a1 [3]int // [0,0,0]

	fmt.Println("a1 short", a1)
	fmt.Println("a1 short %v\n", a1)
	fmt.Println("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool //	[false false false false]
	fmt.Println("a2", a2)

	a3 := [...]int{1, 2, 3, 3, 4}
	fmt.Println("a3", a3)

}
