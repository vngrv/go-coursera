package main 

import "fmt"

const pi = 3.141

const (
  hello = "hello"
  e     = 2.718
)

const (
	zero = iota
	_		 // empty value
	trhee // 3
)

const (
	_ = iota									// empty
	KB uint64 = 1<<(10*iota)	// 1024
	MB												// 1048576
)

const (
	// nontyped const
	year = 2017

	// typed const 
	YearTyped int = 2017

)

func main() {
	var mounth int32 = 13
	// 2030
	fmt.Println(mounth + year)
	// mismatched types int32 & int
	//fmt.Println(mounth + YearTyped)

}

