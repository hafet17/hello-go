package main

import "fmt"

func main() {

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	slice1[0] = "Mei Lagi"

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// data slice tidak akan menambah jika melebihi capacity array
	// var slice2 = month[10:]
	// fmt.Println(slice2)

	// data slice akan menambah jika tidak melebihi capacity array
	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bulan Baru")
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(months)

	//	make slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Hafid"
	newSlice[1] = "Masruri"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSLice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSLice)

}
