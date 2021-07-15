package main

import "fmt"

// function type declarations
type Filter func(string) string

func sayFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" || name == "jangkrik" || name == "cok" {
		return "dilarang mesoh ok ... "
	} else {
		return name
	}
}

func main() {
	sayFilter("alhamdulillah", spamFilter)
	sayFilter("jangkrik", spamFilter)
}
