package main

import "fmt"

func main() {

	var nilaiAkhir = 90
	var absensi = 80

	var nilaiLulus = nilaiAkhir >= 90
	var absensiLulus = absensi >= 80

	var lulus = nilaiLulus && absensiLulus

	fmt.Println(nilaiLulus)
	fmt.Println(absensiLulus)
	fmt.Println(lulus)

}
