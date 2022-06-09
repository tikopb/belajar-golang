package main

import "fmt"

func main() {
	var total = sumAll(1, 2, 3, 1, 2, 3)
	fmt.Println(total)

	numbArray := []int{10, 20, 10, 30}
	fmt.Println(sumAll(numbArray...))
}

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}
