package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	var number1 int = 11
	fmt.Println(number1)

	// Type inference
	var email = "hunghh.dev@gmail.com"
	fmt.Println(email)

	// Khai bao nhieu bien
	// 1. Khai bao nhieu bien cung kieu du lieu
	var a, b int
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)
	var a1, b1 int = 10, 11
	println(a1)
	println(b1)

	var a2, b2 int = 20, 21
	println(a2)
	println(b2)
	// 2. Khai bao nhieu bien khac kieu du lieu
	var (
		name string = "hung"
		address string = "VN"
		age int = 18
	)
	println(name)
	println(address)
	println(age)

	var name1, address1, age1 = "name", "address", 18
	println(name1)
	println(address1)
	println(age1)


	language := "GoLang"
	println(language)

}
