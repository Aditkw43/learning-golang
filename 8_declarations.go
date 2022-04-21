package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpEko NoKTP = "1111111"
	var marriedStatus Married = true
	fmt.Println(ktpEko);
	fmt.Println(marriedStatus);
}