package main

import "fmt"

func main() {
	// Khai bao slice
	var slice []int
	fmt.Println(slice)

	//Khai bao va khoi tao gia tri
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	//Tao slice tu 1array
	var array = [4]int{1, 2, 3, 4}
	slice2 := array[2:4]
	fmt.Println(slice2)
	slice3 := array[:]
	fmt.Println(slice3)

	//Create slice from slice
	slice4 := slice3
	fmt.Println(slice4)
	slice5 := slice4[:3]
	fmt.Println(slice5)

	//Slice is reference type
	var array1 = [5]int{1, 2, 3, 4, 5}
	slice6 := array1[:]
	slice6[0] = 999
	fmt.Println(slice6)
	fmt.Println(array1)

	//length vs capacity of slice
	countries := [...]string{"VN", "USA", "JP", "CN", "RU"}
	slice7 := countries[1:2]
	fmt.Println(slice7)
	fmt.Println(len(slice7))
	fmt.Println(cap(slice7)) // capacity so luong phan tu con lai trong array, tinh tu phan tu dau tien trong array
	fmt.Println("===")
	// make, copy, append
	slice8 := make([]int, 2, 5)
	fmt.Println(slice8)
	fmt.Println(len(slice8))
	fmt.Println(cap(slice8))
	slice9 := make([]int, 2)
	fmt.Println(slice9)
	fmt.Println(len(slice9))
	fmt.Println(cap(slice9))

	slice9 = append(slice8, 100)
	fmt.Println(slice9)
	fmt.Println(len(slice9))
	fmt.Println(cap(slice9))

	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)
	fmt.Println(dest)
	size := copy(dest, src)
	fmt.Println(dest)
	fmt.Println(size)

	fmt.Println("===")
	// delete item with index (trick)
	src = append(src[:1], src[2:]...)
	fmt.Println(src)
}
