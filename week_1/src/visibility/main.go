package main

import (
	"fmt"
	"visibility/person"
)

func main() {
	p := person.NewPerson(1, "vengerov", "sever")

	fmt.Printf("main.PrintPerson: %+v\n", p.secret)

	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
}
