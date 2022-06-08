package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "tiko"
	names[1] = "putra"
	names[2] = "bagaskara"

	var returnvalueString string

	for i := 0; i < len(names); i++ {
		returnvalueString += names[i] + " "
	}

	fmt.Println(returnvalueString)
	names[2] = "testing"
	fmt.Println(names)

	var bilangan []int
	bilangan = append(bilangan, 1)
	bilangan = append(bilangan, 2)
	printSlice(bilangan)
	bilangan = append(bilangan, 3)
	printSlice(bilangan)
}

func printSlice(s []int) { // <- example of declare function
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
