package main

import "fmt"

func main() {

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 == value2) // false
	fmt.Println(value2 > value1)  // true
	fmt.Println(value1 > value2)  // false
	fmt.Println(value1 != value2) // true
	fmt.Println(value1 < value2)  // true
	fmt.Println(value2 < value1)  // false

}
