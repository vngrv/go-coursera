package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"
}

func main() {

	// отложенная задача
	defer fmt.Println("After work")

	// отложенный вызов функции
	defer fmt.Println(getSomeVars())

	// отложенный вызов анонимной функции
	defer func() {
		fmt.Println(getSomeVars())
	}()

	// первое что выполнится
	fmt.Println("Some userful work")
}
