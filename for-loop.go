package main

import "fmt"

func main() {
	counter := 1
	for counter < 10 {
		fmt.Println("perulangan pada ", counter)
		counter++
	}

	slice := []string{"TIKO", "PUTRA", "BAGASKARA"}
	for _, value := range slice {
		fmt.Println(value)
	}
}
