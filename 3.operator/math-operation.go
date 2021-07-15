package main

import "fmt"

func main() {

	a := 10
	b := 10
	c := a * b

	// augmented assignments
	i := 10
	i += 10 // i = i + 10

	// unary op
	i++ // i = i + 1

	fmt.Println(c)
	fmt.Println(i)

}
