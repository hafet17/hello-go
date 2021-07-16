package main

import "fmt"

// looping
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loops := factorialLoop(5)
	fmt.Println(loops)
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
