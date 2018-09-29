package main

import "fmt"

func main() {
	// значение по умолчанию - 0
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// инициализирован тип чисел
	var num2 = 20

	fmt.Println(num0, num1, num2)

	// короткое объявление переменной
	num := 30

	// только для новых переменных
	// no nre varibles on left side of :=
	// num := 31

	num += 1

	fmt.Println("++", num)

	// camelCase - принятый стиль
	userIndex := 10

	// under_score - не принято
	user_index := 10

	fmt.Println(userIndex, user_index)

	var weight, height, age int = 11, 21, 21

	fmt.Println(weight, height, age)

}	