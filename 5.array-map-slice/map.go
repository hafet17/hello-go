package main

import "fmt"

func main() {

	// map = [params1: type key, params2: type value]

	// make 1
	shinobi := map[string]string{
		"name":     "Hayate Gekkou",
		"afliansi": "Konohagakure",
	}

	// add new key
	shinobi["rank"] = "Jounin Tokubetsu"
	fmt.Println(shinobi)
	fmt.Println(shinobi["rank"])

	// make 2
	akatsuki := make(map[string]string)
	akatsuki["member1"] = "Nagato Uzumaki"
	akatsuki["member2"] = "Yahiko"
	akatsuki["member3"] = "Konan"
	akatsuki["member4"] = "Tobi"

	delete(akatsuki, "member4")
	fmt.Println(akatsuki)

}
