package main

import "fmt"

func main() {

	firstName, lastName, _ := getName()
	getName, getAge := getIdentity()

	fmt.Println(firstName, lastName)
	fmt.Println(getName, getAge)
}

func getName() (string, string, string) {
	return "Hafid", "Masruri", "Uchiha"
}

func getIdentity() (string, int) {
	return "Hafid", 22
}
