package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Fatih"
	names[1] = "Nauval"
	names[2] = "Azzidan"

	fmt.Println(names[0])

	var values = [3]int{
		90,
		88,
		77,
	}

	values[0] = 99
	fmt.Println(values)
	fmt.Println("Panjang array names = ", len(names))
	fmt.Println("Panjang array values = ", len(values))

	// step 2
	var memberList = [10]string{
		"hafid",
		"deddy",
		"dian",
		"nopal",
		"fina",
	}

	fmt.Println(memberList)

}
