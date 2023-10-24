package main

import (
	"fmt"
	"time"
)

func main() {

	data := make(chan int)
	QtdeWorkers := 10

	for i := 0; i < QtdeWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}

}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
