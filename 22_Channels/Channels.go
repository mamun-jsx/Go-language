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

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("processing ")

}

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("sending email to ", email)
		time.Sleep(time.Millisecond * 10) // Reduced sleep time so it finishes quickly
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10

	}()
	go func() {
		chan2 <- "pong"
	}()

	for i:=0; i<2; i++{
		select{
			case chan1Value:=<-chan1: 
			fmt.Println("recived chan1 value ", chan1Value)
			case chan2Value:=<-chan2:
			fmt.Println("recived chan2 value ", chan2Value)
		}
		
	}

	// -------------------------------
	// without blocking
	// -------------------------------
	emailChan := make(chan string, 4)
	done := make(chan bool)
	go emailSender(emailChan, done)

	for i := 0; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	// FIX: We must close the channel first!
	// This tells the 'range emailChan' loop in emailSender to terminate.
	close(emailChan)

	// Now we wait for the emailSender goroutine to signal that it has finished.
	<-done

	fmt.Println("done sending")

	// result := make(chan int)
	// go sum(result, 4, 5)
	// showResult := <-result
	// fmt.Println("show result ", showResult)
	// make number channels
	numbChan := make(chan int)
	go processNumb(numbChan)
	// numbChan <- 10

	// for {
	// 	numbChan <- rand.Intn(100)
	// }

	// boolChannels
	// done := make(chan bool)
	// go task(done)
	// <-done // block the program until the task is done

}

// channels are blocking
// when 2nd side not ready to receive data then it will wait for it
// also when 1st side not ready to send data then it will wait for it
