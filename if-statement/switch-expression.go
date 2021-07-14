package main

import "fmt"

func main() {

	name := "hafid"

	switch name {
	case "hafid":
		{
			fmt.Println("halo", name)
		}
	case "nauval":
		{
			fmt.Println("halo", name)
		}
	default:
		fmt.Println("undefined")
	}

	//	switch short statement
	switch length := len(name); length > 5 {
	case true:
		{
			fmt.Println("200 0K")
		}
	case false:
		fmt.Println("404 Error Not Found")
	}
}
