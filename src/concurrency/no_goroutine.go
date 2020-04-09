package main

import (
	"fmt"
	"time"
)

func createPizza(pizza int) {
	time.Sleep(time.Second)
	fmt.Printf("Created Pizza %d \n", pizza)
}

func timeTrack(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
}

func main() {
	defer timeTrack(time.Now(), "Build pizzas")

	for i := 0; i < 10; i++ {
		createPizza(i)
	}
}
