package main

import "fmt"

type Student struct {
	name string
}

//Define method
// func (t Type) methodName(params) return {//body code}
//(t Type) => receiver
//1. Value receiver
//2. Pointer receiver
func (s Student) getName() string {
	return s.name
}

// 1. Value receiver
func (s Student) changeName() {
	fmt.Printf("p2 = %p\n", &s) // no se tao ra 1 gia tri copy, chi dung trong method, khong anh huong field goc
	s.name = "hello"
}

// 2. Pointer receiver
func (s *Student) changeName2() {
	fmt.Printf("p3 = %p\n", s) // anh huong toi field goc
	s.name = "hello"
}

// non-struct
type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s, %s", s, str)
}

func main() {
	student := &Student{"xinchao"}

	name := student.getName()
	fmt.Println(name)

	fmt.Printf("p1 = %p\n", student)
	student.changeName()

	student.changeName2()
	fmt.Println(student.name)
	fmt.Println()
	//non-struct
	s1 := String("z")
	newStr := s1.append("a")
	fmt.Println(newStr)
}
