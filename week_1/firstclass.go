package main

import "fmt"

func doNothing() {
	fmt.Println("I'am regular function")
}

func main() {

	// анонимная функция
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nobody")

	// присвоение анонимной функции в переменную
	printer := func(in string) {
		fmt.Println("printer out:", in)
	}
	printer("as variable")

	// определяем тип функции
	type strFuncType func(string)

	// !!!
	// функция принимает каллбеку
	worker := func(callback, strFuncType) {
		callback("as callback")
	}
	// ./firstclass.go:26:17: undefined: callback
	// ./firstclass.go:27:3: undefined: callback
	// worker(printer)
	//!!!

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefixer, in)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")

}
