package main

import (
	"fmt"
)

type Person struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры, для которой вызыван
func (p Person) UpdateName(name string) {
	p.Name = name
}

// изменяет оригинальную структуру тк ссылается на нее по указателю
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	// pers := Person{1, "Vengerov"}
	pers := new(Person)
	pers.SetName("Vengerov Dmitry")
	// от адреса
	// (&pers).SetName("Vengerov Dmitry")
	// fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "Vengerov",
		Person: Person{
			Id:   2,
			Name: "Met Vengerov",
		},
	}

	acc.SetName("vengerov.met")
	// fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(10)
	fmt.Println(sl.Count(), sl)
}
