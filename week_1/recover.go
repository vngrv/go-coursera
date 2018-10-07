package main

import (
	"fmt"
)

// не упала в панику
func deferTest() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened FIRST:", err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened SECOND:", err)
			panic(2)
		}
	}()

	fmt.Println("some useful work")
	panic("something bad happened")
	return
}

func main() {
	// после паники пришло восстановление
	// никогда не используйте панику и восстановление после паники как эмуляцию блока try-catch.
	deferTest()
	return
}
