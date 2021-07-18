package main

import "fmt"

// normal struct
type Product struct {
	Name, Category string
	price, qty     int
}

func (product Product) sayStruct(name string) {
	fmt.Println("Hello", name, "My Name Is", product.Name)
}

func main() {

	// struct literal
	productList := Product{
		Name:     "",
		Category: "",
		price:    0,
		qty:      0,
	}

	productList.sayStruct("Rujak Cingur")
	fmt.Println(productList)
}
