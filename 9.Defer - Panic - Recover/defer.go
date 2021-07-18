package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	result := 10 / value
	fmt.Println(result)
	fmt.Println("Run Application")
}

func main() {
	runApplication(0)
}
