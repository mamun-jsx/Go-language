# Golang Core Concepts & FAQ

This document organizes and answers the fundamental Go concepts you've been exploring.

---

## 1. Concurrency & Synchronization

### 1.1 What is a Deadlock in Golang?
A **deadlock** occurs when two or more goroutines are waiting for each other to finish, or when a single goroutine is waiting for a channel operation that will never happen (e.g., trying to read from an empty, unclosed channel while no other goroutine is sending data). When Go detects that all goroutines are asleep and unable to wake up, it crashes with a `fatal error: all goroutines are asleep - deadlock!`.

### 1.2 WaitGroup vs. Channels: What is the core difference?
While both are used to synchronize goroutines, they serve different primary purposes:
- **WaitGroup**: Used for **Waiting**. It waits for a collection of goroutines to finish their execution. You don't care about returning data, you just want to know when they are done.
- **Channels**: Used for **Communication**. They are pipes that connect concurrent goroutines, allowing them to send and receive data safely.

### 1.3 What is blocking in channels, and how do Buffers solve this?
- **Unbuffered Channels (Blocking)**: By default, sending data into a channel blocks the sender until another goroutine is ready to receive it. Similarly, receiving blocks until data is sent.
- **Buffered Channels (Non-blocking up to capacity)**: By adding a buffer (`make(chan int, 3)`), you allow the sender to push up to 3 items into the channel without waiting for a receiver. It only blocks when the buffer is completely full.

### 1.4 What is the `select` case?
The `select` statement lets a goroutine wait on multiple channel operations simultaneously. It works like a `switch` statement but for channels. It blocks until one of its cases (a send or receive) is ready. If multiple are ready, it picks one randomly.

---

## 2. Functions & Methods

### 2.1 What is the `defer` function and when to use it?
`defer` delays the execution of a function until the surrounding function returns. 
**Purpose**: It is mainly used to ensure that cleanup tasks are performed, regardless of how the function exits (even if it panics). 
**Use Cases**: Closing files, closing network connections, or unlocking a Mutex (`defer mu.Unlock()`).

### 2.2 What is a Receiver Function in Golang?
Go does not have classes. Instead, it has **Receiver Functions** (or Methods). A receiver function is just a normal function that is bound to a specific struct or type. 
**Example**: `func (p *post) inc()` binds the `inc` method to the `post` struct, allowing you to call `myPost.inc()`.

---

## 3. Data Structures

### 3.1 Array vs. Slice: When to use which?
- **Array**: Fixed length. You must know the exact size at compile time (e.g., `[5]int`). Use arrays only when you strictly need a fixed size (which is rare).
- **Slice**: Dynamic length. A slice is a flexible view into an underlying array (e.g., `[]int`). **Use slices 99% of the time** in Go because they can grow and shrink dynamically.

### 3.2 Slices and the `append` function
Because slices are dynamic, you can add new elements to them using the built-in `append` function (e.g., `mySlice = append(mySlice, 10)`). Under the hood, if the slice runs out of capacity, `append` automatically creates a new, larger array and copies the data over.

---

## 4. Types, Interfaces & Control Flow

### 4.1 How do Structs and Interfaces work together?
- **Structs** hold data (properties/fields).
- **Interfaces** define behavior (a set of method signatures).
If a struct implements all the methods defined in an interface, it automatically satisfies that interface (no `implements` keyword needed). This allows for polymorphism, meaning a function can accept any struct as long as it satisfies the required interface.

### 4.2 Generics in Golang
Introduced in Go 1.18, Generics allow you to write functions or data structures that can work with any data type, rather than just one specific type. You use type parameters (like `[T any]`) so you don't have to write the same function multiple times for `int`, `string`, `float64`, etc.

### 4.3 `for` loop vs. `range`: When to use which?
- **Standard `for` loop** (`for i := 0; i < 10; i++`): Use this when you need a counter, when you are not looping over a collection, or when you need to skip specific indexes manually.
- **`range` loop** (`for index, value := range mySlice`): Use this when you are iterating over the contents of a Slice, Array, Map, or Channel. It is cleaner, safer, and less prone to "index out of bounds" errors.