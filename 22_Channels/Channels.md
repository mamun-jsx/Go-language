# Understanding Channels in Go

Channels in Go are like pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine. They provide a safe way for goroutines to communicate and synchronize their execution.

## The `Channels.go` Code Explained

### Code Overview

The provided code demonstrates a basic use case of channels where a main goroutine generates random numbers and sends them to a background goroutine for processing.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// processNumb receives data from the channel and processes it
func processNumb(numbChan chan int) {
	// The range loop continuously reads from the channel until it is closed
	for num := range numbChan {
		fmt.Println("processing number ", num)
		time.Sleep(time.Second * 1) // Simulate work
	}
}

func main() {
	// 1. Create a new channel of type int
	numbChan := make(chan int)
	
	// 2. Start the processing goroutine
	go processNumb(numbChan)
	
	// 3. Continuously send random numbers into the channel
	for {
		numbChan <- rand.Intn(100) // Sends a random integer between 0-99
	}
}
```

### Key Concepts from the Code
1. **Creation**: `make(chan int)` creates an unbuffered channel that can transmit integers.
2. **Sending Data**: `numbChan <- rand.Intn(100)` sends data *into* the channel.
3. **Receiving Data**: `for num := range numbChan` loops indefinitely, receiving data *from* the channel as soon as it's available.
4. **Blocking Nature**: Channels block by default. 
   - When the main goroutine sends data (`numbChan <-`), it blocks until `processNumb` is ready to receive it.
   - Conversely, `processNumb` blocks on `range numbChan` until the main goroutine sends something.

### Output
When running the code, it continuously outputs a random number every second:
```text
processing number  47
processing number  75
processing number  9
... (runs indefinitely until interrupted with Ctrl+C)
```

---

## Example: Odd / Even Checker with Channels

Here is an example of using a channel to check if numbers are odd or even, and returning a structured result through a channel.

### The Code

```go
package main

import (
	"fmt"
)

// Result struct to hold both the number and its odd/even status
type Result struct {
	Number int
	IsEven bool
}

// checkOddEven receives numbers, determines if they are even/odd, and sends results back
func checkOddEven(inputChan chan int, outputChan chan Result) {
	for num := range inputChan {
		// Determine if the number is even
		isEven := num%2 == 0
		
		// Send the structured result to the output channel
		outputChan <- Result{Number: num, IsEven: isEven}
	}
	// Close the output channel when we are done processing all inputs
	close(outputChan)
}

func main() {
	// Create channels
	inputChan := make(chan int)
	outputChan := make(chan Result)

	// Start the worker goroutine
	go checkOddEven(inputChan, outputChan)

	// Send numbers to the input channel in a separate goroutine
	go func() {
		numbersToCheck := []int{1, 2, 3, 4, 5, 10, 15}
		for _, n := range numbersToCheck {
			inputChan <- n
		}
		// Close the input channel to signal we are done sending data
		close(inputChan)
	}()

	// Read and print the results from the output channel
	for res := range outputChan {
		if res.IsEven {
			fmt.Printf("%d is Even\n", res.Number)
		} else {
			fmt.Printf("%d is Odd\n", res.Number)
		}
	}
}
```

### How it Works
1. **Structs over Channels**: We define a `Result` struct to bundle the original number and the boolean result together. The `outputChan` is of type `chan Result`.
2. **Multiple Goroutines**: We have a goroutine that generates numbers, a worker goroutine (`checkOddEven`) that processes them, and the `main` goroutine that prints the results.
3. **Closing Channels**: We use `close(inputChan)` to let the `checkOddEven` function know no more numbers are coming. In turn, `checkOddEven` closes the `outputChan` when it finishes, which allows the `range` loop in `main` to terminate naturally instead of causing a deadlock.

### Output
```text
1 is Odd
2 is Even
3 is Odd
4 is Even
5 is Odd
10 is Even
15 is Odd
```
