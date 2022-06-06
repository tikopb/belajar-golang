package main

import (
	"fmt"
)

func main() {

	type NoKtp string //-> declaration alias menggunakan type alih alih menggunakan var
	type Married bool

	var NoKtpTik NoKtp = "1029109238109842"
	var marriedStatus Married = true
	fmt.Println(NoKtpTik)
	fmt.Println(marriedStatus)
}
