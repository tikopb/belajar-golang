package main

import "fmt"

func main() {
	name := "TIKO"

	switch name {
	case "TIKO":
		fmt.Println("Hello Tiko")
	case "PUTRA":
		fmt.Println("Hello Putra")
	default:
		fmt.Println("not found")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama pas")
	case false:
		fmt.Println("nama terlalu panjang")
	}
}
