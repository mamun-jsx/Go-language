package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// channels is like a pipe connecting two goroutines
// one goroutine send data to the pipe and the other goroutine
// receive data from the pipe

// sending data
func processNumb(numbChan chan int) {
	for num := range numbChan {
		fmt.Println("processing number ", num)
		time.Sleep(time.Second * 1)

	}
}
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult

}

func main() {

	result := make(chan int)
	go sum(result, 4, 5)
	showResult := <-result
	fmt.Println("show result ", showResult)
	// make number channels
	numbChan := make(chan int)
	go processNumb(numbChan)
	// numbChan <- 10

	// for {
	// 	numbChan <- rand.Intn(100)
	// }
}

// channels are blocking
// when 2nd side not ready to receive data then it will wait for it
// also when 1st side not ready to send data then it will wait for it
