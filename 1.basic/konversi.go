package main

import "fmt"

func main() {

	var nilai32 int32 = 7890
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "mohamad hafid"

	var convertByte byte = name[0]

	var convertString string = string(convertByte)
	fmt.Println(convertString)

}
