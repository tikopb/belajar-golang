package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "tiko",
		"address": "semarang",
	}

	fmt.Println(person)
	fmt.Println(person["address"])

	person["address"] = "surabaya"
	fmt.Println(person)

	delete(person, "address")
	fmt.Println(person)
}
