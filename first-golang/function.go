package main

import "fmt"

// func func_name(params) return_type {//code}

func xinChao()  {
	fmt.Println("Xin Chao")
}

func sayHello(name string)  {
	fmt.Println("Hello ", name)
}

func greeting(name string) string {
	return name
}

// Multiple return values
func rectInfo(w, h int) (int, int) {
	return w, h
}

// Named return values
func rectInfoIs(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	xinChao()
	sayHello("hihi")
	fmt.Println(greeting("hunghh.dev"))
	fmt.Println(rectInfo(6,9))

	w, h, isSquare := rectInfoIs(200, 200)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println(w)
		fmt.Println(h)
	}
}
