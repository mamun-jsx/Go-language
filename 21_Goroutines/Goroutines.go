package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int) {
	fmt.Println("doing task", id)
}
func job(ids int) {
	fmt.Println("doing job ", ids)
}

// serial receives a pointer to a sync.WaitGroup.
// It's important to pass a pointer so that the goroutine modifies the original WaitGroup,
// not a copy. If passed by value, a copy would be made, and calling Done() wouldn't
// decrement the counter in the main function's WaitGroup, leading to a deadlock.
func serial(num int, w *sync.WaitGroup) { // send the pointer to access the wait group from the main
	defer w.Done() // decrease the wait group counter. Defer ensures it's called even if a panic occurs.
	fmt.Println("printing serial ", num)
}
func main() {
	for i := 1; i <= 10; i++ {
		go task(i)
	}

	// do another way or syntex
	for i := 30; i <= 40; i++ {
		go func(ids int) {
			fmt.Println("doing job", ids)
		}(i)
	}
	fmt.Println("------------------------------------")

	// A sync.WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of goroutines to wait for.
	// Then each of the goroutines runs and calls Done when finished.
	// At the same time, Wait can be used to block until all goroutines have finished.
	var wg sync.WaitGroup // create or declare the wait group
	for i := 0; i <= 10; i++ {
		wg.Add(1) // increment the WaitGroup counter for each goroutine launched
		go serial(i, &wg) // Pass a pointer to the WaitGroup to the goroutine
	}
	
	// Wait blocks the calling goroutine (main here) until the WaitGroup counter reaches zero.
	// Without this, the main function would exit before the goroutines finish executing.
	wg.Wait()
	// time.Sleep(time.Second * 2)
}
