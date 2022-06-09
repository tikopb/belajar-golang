package main

import "fmt"

func main() {

	test1, test2, _ := returnString("test1", "test2", "test3")
	fmt.Println(test1 + " " + test2)
}

func returnString(valueBack string, valueBack2 string, valueBack3 string) (string, string, string) {
	return valueBack, valueBack2, valueBack3
}
