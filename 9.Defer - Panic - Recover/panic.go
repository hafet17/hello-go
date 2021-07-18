package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error message:", message)
	}
	fmt.Println("end application")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("error application")
	} else {
		fmt.Println("application is running")
	}
}

func main() {
	runApp(true)
}
