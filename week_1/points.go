package main 

import "fmt"

func main() {
  a := 2
  b := &a
  *b = 3    // a = 3 
	c := &a 	// new pointer for a 

	// get pointer for int val
	
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c and a not changes 

	c = d 
	*c = 14
}