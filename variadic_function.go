package main

import "fmt"

func addItem(item int, list ...int) {
	// list is a slice
	list = append(list, item)
	fmt.Println(list)
}

func changeSlice(list ...int) {
	list[0] = 999
}

func main() {
	addItem(1, 100, 200, 300, 400)

	var listSlice = []int{1, 2, 3, 4}
	// pass a slice
	addItem(100, listSlice...)

	changeSlice(listSlice...)
	fmt.Println(listSlice)
}
