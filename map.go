package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	fmt.Println(myMap)
	fmt.Println(myMap == nil)

	var myMap1 map[string]int
	fmt.Println(myMap1)
	fmt.Println(myMap1 == nil)

	// Khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println(myMap2)

	// Them 1 phan tu vao map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	myMap2["key1"] = 10
	fmt.Println(myMap2)

	// Xoa 1 phan tu
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	// length of map
	fmt.Println(len(myMap2))

	// Map is reference type
	myMap3 := myMap2
	fmt.Println(myMap3)
	myMap2["key5"] = 1000
	fmt.Println(myMap3)
	myMap3["key4"] = 1001
	fmt.Println(myMap2)

	// Truy cap 1 phan tu trong map
	value := myMap2["key2"]
	fmt.Println(value)
	value1, found := myMap2["key200"]
	fmt.Println(found)
	fmt.Println(value1)

}
