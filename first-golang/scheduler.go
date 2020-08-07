package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func g1() {
	fmt.Println("g1")
	os.OpenFile()
}

func g2() {
	fmt.Println("g2")
}

func g3() {
	fmt.Println("g3")
}

func g4() {
	fmt.Println("g4")
}

func main() {
	runtime.GOMAXPROCS(2)
	numberP := runtime.NumCPU()
	fmt.Println(numberP)
	fmt.Println(runtime.GOMAXPROCS(0))

	go g1()
	go g2()
	go g3()
	go g4()
	time.Sleep(time.Second)
}
