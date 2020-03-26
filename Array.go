package main

import (
	"fmt"
)

func main() {
	// Khai bao
	var myArray [4]int
	fmt.Println(myArray)

	// Set value
	myArray[1] = 99
	fmt.Println(myArray)

	// Khai bao + khoi tao
	arrays := [3]int{1, 2, 3}
	fmt.Println(arrays)

	// Size
	fmt.Println(len(myArray))

	// Khai bao khong set size
	arrPro := [...]string{"honda", "yamaha", "suzuki"}
	fmt.Println(arrPro)
	fmt.Println(len(arrPro))

	// array is value type, not is reference type
	countries := [...]string{"VN", "US", "FRANCE"}
	copyCountries := countries
	fmt.Println(countries)
	fmt.Println(copyCountries)
	countries[0] = "JP"
	fmt.Println(countries)
	fmt.Println(copyCountries)
	// => tom lai no se tao ra vung nho moi, OK

	// loop array
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
	for index, value := range countries {
		fmt.Printf("i=%d value=%s\n", index, value)
	}
	for _, value := range countries {
		//bo qua index va nguoc lai neu thich
		// _ is blank identifier
		fmt.Printf("value=%s\n", value)
	}

	// mang hai chieu, thich bao nhieu chieu thi them vao [] :D
	matrix := [4][2]int{
		{1, 2},
		{3, 2},
		{5, 2},
		{7, 2},
	}
	fmt.Println(matrix)
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
