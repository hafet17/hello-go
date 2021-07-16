package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name); len(name) > 5 {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome")
	}
}

func main() {

	result := func(name string) bool {
		return name == "blocked"
	}

	registerUser("hafid", result)
	registerUser("blocked", result)
}
