package main

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	myID := UserID(idx)
	println(uid, myID)

}