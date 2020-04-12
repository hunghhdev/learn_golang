package main

import (
	"fmt"
)

func main() {
	// Unbeffered channel
	ch := make(chan int)

	go func() {
		ch <- 100
		fmt.Println("sent")
	}()
	fmt.Println(<-ch)
	fmt.Println("done")
	fmt.Println()

	// Buffered channel
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2

	close(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println()

	// Select
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}

		done <- true
	}()

	// for {
	// 	select {
	// 	case v := <-queue:
	// 		fmt.Println(v)
	// 	case <-done:
	// 		fmt.Println("done")
	// 		return
	// 	}
	// }
	fmt.Println("---------")

	// Close channel
	queue1 := make(chan int, 10)
	go func() {
		for i := 10; i < 20; i++ {
			if i > 15 {
				close(queue1)
				break
			} else {
				queue1 <- i
			}
		}
	}()

	for value := range queue1 {
		fmt.Println(value)
	}

}
