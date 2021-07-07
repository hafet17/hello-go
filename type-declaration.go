package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpHafid NoKTP = "18929203030303003"
	var marriedStatus Married = false

	fmt.Println(noKtpHafid)
	fmt.Println(marriedStatus)
}
