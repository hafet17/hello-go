package main

import "fmt"

func main() {

	fak := "ai"

	if fak == "teknik" {
		prodi := "informatika"
		if prodi == "informatika" {
			fmt.Println(prodi)
		}
	} else if fak == "ai" {
		prodi := "pai"
		if prodi == "pai" {
			fmt.Println(prodi)
		}
	} else {
		fmt.Println(false)
	}

}
