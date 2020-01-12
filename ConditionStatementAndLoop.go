package main

import "fmt"

func main() {
	// if else
	number := 10
	if number == 10 {
		fmt.Println("number = 10")
	} else if number < 10 {
		fmt.Println("number < 10")
	} else {
		fmt.Println("false")
	}

	// if statement; condition {//code}
	if a := 100; a == 100 {
		fmt.Println("a=100")
	} else {
		fmt.Println("a!=100")
	}

	// switch - case
	vSwitch := 6
	switch vSwitch {
	case 1, 6, 9:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 69:
		fmt.Println("69")
	default:
		fmt.Println("none")
	}

	switch {
	case vSwitch > 10:
		fmt.Println("number > 10")
	case vSwitch < 10:
		fmt.Println("number == 10")
	}

	// FallThrough - Không break mà tiếp tục chạy
	switch number {
	case 1:
		fmt.Println("number = 1")
		fallthrough
	case 10:
		fmt.Println("number = 10")
		fallthrough
	case 2:
		fmt.Println("number = 2")
		fallthrough
	case 3:
		fmt.Println("number = 3")
	}

	// break, goto
	switch number {
	case 1:
		fmt.Println("number = 1")
		fallthrough
	case 10:
		if number == 10 {
			fmt.Println("break here")
			break
		}
		fmt.Println("number = 10")
		fallthrough
	case 2:
		if number == 10 {
			fmt.Println("break here")
			goto handleNumberEqual10
		}
		fmt.Println("number = 2")
		handleNumberEqual10:
			fmt.Println("goto here")
		fallthrough
	case 3:
		fmt.Println("number = 3")
	}



	// Loops
	for i := 1;i<=10 ;i++  {
		fmt.Println(i)
	}
	//* break, continue
	for i := 1;i<=10 ;i++  {
		if i==4{
			continue
		}
		if i==6 {
			break
		}
		fmt.Println(i)
	}

	// while
	for i := 1; i <= 10; {
		fmt.Println(i)
		i++
	}

	// infinite loop
	for {
		fmt.Println("infinite loop")
	}

	// for init, condition, post
	for i, j :=1,2; i<10 && j < 10 ; i,j = i+1, j+1  {
		fmt.Println(i)
		fmt.Println(j)
	}
}
