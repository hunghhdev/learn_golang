package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Bool
	var myBool bool = true // true or false
	println(myBool)

	// String
	var myString string = "Hello World"
	println(myString)

	// int
	var myInt int = 69
	println(myInt)

	// int8, int16, int32, int64
	// 1. Range
	// 8: -128 -> 127
	println(math.MinInt8)
	println(math.MaxInt8)
	// 16: -32768 -> 32767
	println(math.MinInt16)
	println(math.MaxInt16)
	// 32: -2147483648 -> 2147483647
	println(math.MinInt32)
	println(math.MaxInt32)
	// 64: -9223372036854775808 -> 9223372036854775807
	println(math.MinInt64)
	println(math.MaxInt64)
	// 2. Bits
	println(bits.OnesCount8(math.MaxUint8))
	println(bits.OnesCount16(math.MaxUint16))
	println(bits.OnesCount32(math.MaxUint32))
	println(bits.OnesCount64(math.MaxUint64))

	// uint, uint8, uint16, uint32, uint64 - Số nguyên dương
	var myUInt uint = 10
	println(myUInt)

	// Byte
	var myByte byte = 1 // alias for uint8
	println(myByte)
	fmt.Printf("%T\n", myByte)
	var a byte = 'A'
	println(a)
	fmt.Printf("%X", a)
	println()
	// float32, float64
	var myFloat float32 = 1.1
	fmt.Println(myFloat)

	// complex64, complex128 z = a + bi
	var z complex64 = 10+0i
	var z1 complex64 = 10+2i
	fmt.Println(z+z1)

	// Rune
	var myS string = "Nhẫn"
	myRunes := []rune(myS)
	for i:=0;i<len(myRunes) ;i++  {
		fmt.Printf("%c", myRunes[i])
	}
	println()
	var myRune rune = 'Ấ' // alias for int32
	fmt.Printf("%T - %c - %U", myRune, myRune, myRune)
	println()

	// uintptr
	//* Sử dụng lưu địa chỉ của con trỏ


	/// Type conversions
	i := 69
	f := float32(i)
	fmt.Println(f)
	//* Tính toán phải cùng kiểu dữ liệu

	// Constants
	const PI = 3.14
	fmt.Println(PI)
	//* constant cannot be declared using := syntax
}
