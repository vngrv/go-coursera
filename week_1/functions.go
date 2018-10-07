package main

import "fmt"

// один параметр
func singleIn(in int) int {
	return in
}

// несколько параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// именованный результат
func namedReturn() (out int) {
	out = 2
	return
}

// несколько результатов
func multiplereturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happened")
	}
	return in, nil
}

func mulripleNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("some error happened")
		// аналогично return res, err
		// или return 1, fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

// не фиксированное количество параметров
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	// fmt.Println(mulripleNamedReturn(true))
	// return

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return
}
