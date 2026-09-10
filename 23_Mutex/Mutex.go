package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	
	// Why Mutex?
	// When multiple goroutines try to change the same resource (like 'views' here),
	// they can overwrite each other, causing a "Race Condition".
	// Mutex solves this by locking the resource so only ONE goroutine can change it at a time.
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	// defer ensures that even if something fails, the lock is released
	// and the wait group is notified that this goroutine is done.
	defer func() {
		p.mu.Unlock() // Step 3: Unlock the resource for the next goroutine to use
		wg.Done()     // Step 4: Tell WaitGroup this task is finished
	}()
	
	p.mu.Lock() // Step 1: Lock the resource before making any changes
	p.views++   // Step 2: Safely increase the view count
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	// We are launching 100 goroutines at the same time
	for i := 0; i < 100; i++ {
		wg.Add(1)          // Tell WaitGroup we are starting a new goroutine
		go myPost.inc(&wg) // Run the function concurrently
	}
	
	wg.Wait() // Block here until all 100 goroutines have finished

	// The result will perfectly be 100 because Mutex protected the data
	fmt.Println("number of views ", myPost.views)
}
