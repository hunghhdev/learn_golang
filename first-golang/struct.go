package main

import "fmt"

// struct la kieu du lieu tu dinh nghia, duoc cau tao bang tap cac field tuong ung voi cac tap du lieu cu the cho field do

type Student struct {
	id   int
	name string
	info StudentInfo
}

type StudentInfo struct {
	class string
	email string
	age   int
}

func main() {
	// named
	// st1 := Student{
	// 	id:   123,
	// 	name: "hung",
	// }
	// fmt.Println(st1)
	// fmt.Println(st1.id)
	// fmt.Println(st1.name)

	// st2 := Student{456, "hoang"}
	// fmt.Println(st2)
	// fmt.Println(st2.id)
	// fmt.Println(st2.name)

	// var st3 Student = struct {
	// 	id   int
	// 	name string
	// }{
	// 	789,
	// 	"hung",
	// }
	// fmt.Println(st3)
	// fmt.Println(st3.id)
	// fmt.Println(st3.name)

	// anonymous struct
	var anonymous = struct {
		email string
		age   int
	}{
		"gmail",
		19,
	}
	fmt.Println(anonymous)
	fmt.Println(anonymous.email)
	fmt.Println(anonymous.age)

	// anonymous fields
	type NoName struct {
		string
		int
	}
	n := NoName{"van a", 12}
	fmt.Println(n)

	// pointer in struct
	// pointer := &Student{
	// 	789,
	// 	"hunghh",
	// }
	// fmt.Println(pointer)
	// fmt.Println(&pointer)
	// fmt.Println((*pointer).name)
	// fmt.Println(pointer.id)

	// struct in struct - nested struct
	student := Student{
		id:   123,
		name: "xx",
		info: StudentInfo{
			class: "xz",
			email: "dev@hung.com",
			age:   27,
		},
	}
	fmt.Println(student)

	student2 := Student{
		123,
		"xx",
		StudentInfo{
			class: "xz",
			email: "dev@hung.com",
			age:   27,
		},
	}
	fmt.Println(student2)

	//Compare 2 struct, chi compare voi kieu du lieu so sanh duoc
	type struct1 struct {
		id   int
		name string
	}

	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}
	fmt.Println(s1 == s2)

	// Zero value - truong nao khong duoc khoi tao se lay zero value cua truong do
	var sx Student
	fmt.Println(sx)
}
