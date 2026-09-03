# Pointers in Go (Golang)

## What is a Pointer?
A pointer is a variable that stores the **memory address** of another variable. Instead of holding a data value directly (like an integer or string), it holds the location in the computer's memory where that data is stored.

Pointers are very important for:
1. **Low-level memory management**: They allow you to refer to the exact location of data.
2. **Performance**: Passing a memory address (pointer) to a function is much lighter and faster than copying large amounts of data.
3. **Modifying Original Values**: If you pass a variable to a function by value, Go creates a copy. If you want a function to modify the original variable, you must pass a pointer (reference).

## Key Operators
- `&` (Ampersand - Address Operator): Used to get the **memory address** of a variable (called referencing). For example, `&num` gives the memory location of `num`.
- `*` (Asterisk - Pointer/Dereference Operator): 
  - Used in a type declaration to indicate a pointer (e.g., `*int` means "pointer to an int").
  - Used on a pointer variable to access or change the actual value stored at that memory address (called dereferencing). For example, `*numbe = 15`.

---

## Example Code

This example demonstrates the difference between **Pass by Value** (which doesn't change the original variable) and **Pass by Reference** using pointers (which successfully modifies the original variable).

```go
package main

import "fmt"

// 1. Pass by Value
// This function receives a COPY of the value.
// Changes made here will NOT affect the original variable in main().
func changeNum(num int) {
	num = 5
	fmt.Println("Inside changeNum (copy modified):", num)
}

// 2. Pass by Reference (Using Pointer)
// This function receives the MEMORY ADDRESS of the value.
// Changes made here WILL affect the original variable in main().
func changAbleNumber(numbe *int) { // *int means it expects a pointer to an int
	
	// *numbe is dereferencing: we go to the memory address and change the actual value to 15
	*numbe = 15 

	fmt.Println("Inside changAbleNumber (memory address received):", numbe)
}

func main() {
	// === Pass by Value Example ===
	num := 1
	fmt.Println("Original 'num' before changeNum:", num)
	
	changeNum(num) // Sending a copy of 'num'
	
	fmt.Println("Original 'num' after changeNum:", num) 
	// As you can see, 'num' remains 1 in main.

	fmt.Println("Memory location of 'num':", &num) // Getting the address using &
	
	fmt.Println("--------------------------------------------------")

	// === Pass by Reference Example ===
	toChangeNumbe := 100
	fmt.Println("Original 'toChangeNumbe' before changAbleNumber:", toChangeNumbe)
	
	// We pass the memory address of toChangeNumbe using &
	changAbleNumber(&toChangeNumbe) 
	
	fmt.Println("Original 'toChangeNumbe' after changAbleNumber:", toChangeNumbe) 
	// The original value has successfully been changed to 15!
}
```

## Output

```text
Original 'num' before changeNum: 1
Inside changeNum (copy modified): 5
Original 'num' after changeNum: 1
Memory location of 'num': 0x1b1eebf08020
--------------------------------------------------
Original 'toChangeNumbe' before changAbleNumber: 100
Inside changAbleNumber (memory address received): 0x1b1eebf08028
Original 'toChangeNumbe' after changAbleNumber: 15
```

### Summary of what happened in the output:
1. In `changeNum`, we changed `num` to `5`, but it only changed the local copy inside that function. Back in `main`, it remained `1`.
2. In `changAbleNumber`, we passed the exact memory address (`&toChangeNumbe`). So when we did `*numbe = 15`, it went straight to that address and updated the original variable. Back in `main`, the variable was successfully updated to `15`.
