package main

import "fmt"

// Con tro la 1bien, no luu tru vung nho cua bien khac
func main() {
	a := 100.100

	var pointer *float64
	pointer = &a
	fmt.Println(pointer)
	fmt.Printf("%T", pointer)
	fmt.Println("\n===========")

	b := 123
	p2 := new(int)
	p2 = &b
	fmt.Println(p2)
	fmt.Printf("%T", p2)
	fmt.Println("\n===========")

	// zero value
	var p3 *int
	fmt.Println(p3)
	p4 := new(int)
	fmt.Println(p4)
	fmt.Println("\n===========")

	c := 1000
	var p5 *int
	p5 = &c
	fmt.Println(p5)
	*p5 = 999 // c = 999
	fmt.Println(c)
	fmt.Println(p5)
	fmt.Println(*p5)
	fmt.Println("\n===========")

	// pointer in array
	array := [...]int{1, 2, 3}
	var p6 *[len(array)]int
	p6 = &array
	fmt.Println(p6)
	fmt.Printf("%T", p6)
	fmt.Println("\n===========")

	// pointer call function
	d := 30
	var p7 *int = &d
	applyPointer(p7)
	fmt.Println(d)
}

func applyPointer(pointer *int) {
	fmt.Println(pointer)
	*pointer = 777
}
