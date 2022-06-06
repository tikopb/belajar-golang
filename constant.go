package main

import "fmt"

func main() {
	const (
		firstName string = "TiKO"
		lastName  string = "Putra"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	var e byte = firstName[0]
	var eString string = string(e)
	fmt.Println(eString)
}
