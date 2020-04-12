package main

import "fmt"

// Animal interface
type Animal interface {
	speak()
}

// Movement interface
// Multiple interfaces
type Movement interface {
	move()
}

// NextAnimal interface
// Embed interface
type NextAnimal interface {
	Movement
	Animal
}

// Dog dog
type Dog struct {
}

func (d Dog) speak() {
	fmt.Println("haha")
}

func (d Dog) move() {
	fmt.Println("run")
}

// Empty interface
func goout(i interface{}) {
	fmt.Println(i)
}

func main() {
	var animal Animal
	animal = Dog{}
	animal.speak()

	dog := Dog{}
	var m Movement = dog
	m.move()
	var a Animal = dog
	a.speak()

	fmt.Println("")

	var na NextAnimal = dog
	na.move()
	na.speak()

	goout("x")
	goout(123)
}
