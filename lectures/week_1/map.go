package main

import "fmt"

func main() {
	// init map
	var user map[string]string = map[string]string{
		"name":     "Met",
		"lastname": "Vengerov",
	}

	// copy
	profile := make(map[string]string, 10)
	mapLength := len(user)
	fmt.Printf("%d %+v\n", profile, mapLength)

	// если ключа нет - вернее значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	// удаление ключа
	delete(user, "lastName")
	fmt.Println("%#v\n", user)
}
