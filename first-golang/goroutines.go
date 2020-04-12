package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func g1() {
	fmt.Println("g1")
}

func g2() {
	fmt.Println("g2")
	wg.Done()
}

func g3() {
	fmt.Println("g3")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	go g1()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)

	// Synchronized goroutines
	fmt.Println("begin")
	wg.Add(2)
	go g2()
	go g3()
	wg.Wait()
	fmt.Println("end")

	/*
		logic 1
		go g1()
		go g2()

		logic 2

	*/
}
