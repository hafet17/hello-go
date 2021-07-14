package main

import "fmt"

func main() {

	// deklarasi variabel
	var nama string
	var umur int

	nama = "Mohamad Hafid Masruri"
	fmt.Println("nama = ", nama)

	umur = 22
	fmt.Println("umur = ", umur)

	// mempersingkat variabel
	var alamat = "Konoha"
	fmt.Println("alamat = ", alamat)

	// lebih singkat
	negara := "Indonesia"
	fmt.Println("negara = ", negara)

	// multiple variabel
	var (
		firstName = "Nauval"
		lastName  = "Azzidan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
