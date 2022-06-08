package main

import "fmt"

func main() {
	var name1 = "joko"
	var name2 = "joko"

	var validation bool = name1 == name2 // compare ini berjalan secara case sensitive

	fmt.Println(validation)

	var angka1 = 12
	var angka2 = 14

	fmt.Println(angka1 < angka2)

}
