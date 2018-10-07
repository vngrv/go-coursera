package main

import (
	"fmt"
)

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	// полное объявление структуры
	var acc Account = Account{
		Id:   1,
		Name: "vengerov",
		Person: Person{
			Name:    "Metthew",
			Address: "Moskow",
		},
	}
	fmt.Printf("%#v\n", acc)
	// короткое объявление структуры
	acc.Owner = Person{2, "Vengerov Met", "Moskow"}

	fmt.Printf("%#v\n", acc.Name)
	fmt.Printf("%#v\n", acc.Person.Name)
}
