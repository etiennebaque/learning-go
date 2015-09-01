package main

// Concurrency: program that runs 4 routines, as a test to experiment concurrency using channels.

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker model
type Worker struct {
	Id            int
	NumOperations int
}

// Calculate method
func (w *Worker) calculate(x, y chan int) {
	for {
		x_data := <-x
		y_data := <-y
		w.NumOperations += 1
		fmt.Printf("Worker %d calculated %d + %d\n", w.Id, x_data, y_data)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	firstNum, secondNum := make(chan int), make(chan int)
	for i := 0; i < 4; i++ {
		worker := &Worker{Id: i}
		go worker.calculate(firstNum, secondNum) // Goroutine, channels as a params of calculate method
	}

	for x := 0; x < 20; x++ {
		firstNum <- rand.Intn(10000)
		secondNum <- rand.Intn(10000)
	}
	fmt.Println("All done")
}
