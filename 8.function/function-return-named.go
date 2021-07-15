package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Mohamad"
	middleName = "Hafid"
	lastName = "Masruri"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
