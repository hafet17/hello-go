package main

import "fmt"

func main() {

	counter := 0
	name := "Hafid"

	increment := func() {
		fmt.Println("increment")
		counter++
		name = "Budi"
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)

}
