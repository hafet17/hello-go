package main

import (
	"fmt"
	"strings"
)

// 1. function normal
func printMsg(name string, names []string) {
	join := strings.Join(names, "")
	fmt.Println(name, join)
}
func testArray(arr [5]int) {
	fmt.Println(arr)
}

// 2. return value
func testValue(value1, value2 int) int {
	return value1 * value2
}

// 3. function as parameter - type declarations
type filterName func(name string) string
type filterValue func(value int) string

func helloFilter(name string, value int, filName filterName, filValue filterValue) {
	nameFiltered := filName(name)
	valueFiltered := filValue(value)

	fmt.Println(nameFiltered)
	fmt.Println(valueFiltered)
}

// function params 1
func helloName(name string) string {
	if name == "hafid" {
		return "login succesfully"
	} else {
		return "invalid login"
	}
}

// function params 2
func helloValue(value int) string {
	if value >= 70 {
		return "202 OK"
	} else {
		return "404 Not Found"
	}
}

// 4. Multiple Value
func helloMultiple() (string, string, string) {
	return "Fatih", "Nauval", "Azzidan"
}

func main() {

	// result 1
	names := []string{"John", "Wick"}
	printMsg("hello", names)
	testArray([5]int{1, 3, 4, 5, 6})

	// result 2
	tes := testValue(10, 10)
	fmt.Println(tes)

	//	result 3
	helloFilter("hafid", 90, helloName, helloValue)

	// result 4
	firstName, middleName, _ := helloMultiple()
	fmt.Println(firstName, middleName)

}
