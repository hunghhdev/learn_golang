package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func createPizza(pizza int) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Created Pizza %d \n", pizza)
}

func timeTrack(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
}

func main() {
	defer timeTrack(time.Now(), "Build Pizzas")

	var numberOfOvens int = 100

	runtime.GOMAXPROCS(numberOfOvens)
	wg.Add(numberOfOvens)

	for i := 0; i < numberOfOvens; i++ {
		go createPizza(i)
	}
	wg.Wait()

}
