package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for statement
	// init statement, kondisi , post statement
	for loop := 1; loop <= 10; loop++ {
		fmt.Println("perulangan ke", loop)
	}

	// for iteration
	slice := []string{"Hafid", "Masruri", "Nauval", "Azzidan"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range
	array := [...]string{"hafid", "nia", "shinta", "ina"}
	for _, value := range array {
		//fmt.Println("index ke", i , "=",value)
		fmt.Println(value)
	}

	person := map[string]string{
		"name":  "hafid",
		"hobby": "ngoding",
	}
	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
