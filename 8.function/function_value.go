package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye" + name
}

func main() {
	goodBye := getGoodBye // function set value
	fmt.Println(goodBye("hafid"))
}
