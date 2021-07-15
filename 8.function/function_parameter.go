package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

func sayArray(arr [2]string) {
	fmt.Println(arr)
}

func sayMap(person map[string]string) {
	fmt.Println(person)
}

func main() {
	sayHello("Hafid", "Masruri")
	sayArray([2]string{"hafid", "masruri"})
	sayMap(map[string]string{
		"name":    "hafid",
		"address": "konoha",
	})
}
