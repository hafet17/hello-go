package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"John", "Wick"}
	printMessage("halo", names)
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
