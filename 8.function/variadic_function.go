package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumKali(numbers ...int) int {
	total := 1
	for _, number := range numbers {
		total *= number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 90, 80, 40)
	fmt.Println(total)

	kali := sumKali(5, 10, 2, 4, 2)
	fmt.Println(kali)

	slice := []int{10, 80, 70, 30, 20}
	fmt.Println(sumAll(slice...))
}
