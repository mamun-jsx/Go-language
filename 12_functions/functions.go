package main

import "fmt"

// function to add some int

func add(a int, b int) int {
	return a + b
}

// function to add multiple int
// * if value is similar then we can use type one time
func mul(c, d int) int {
	return c * d
}

// a cool feature into go lang that can return miltiple value
func getMultipleData() (string, int, bool, string, any) {
	return "Golang", 43, true, "C++", nil
}

// go func can be asign into a variable also and one function can be pass to another as argument, we can return a function into another function

// ============
func processIt(fn func(k int) int) {
	fn(2)
}

func main() {
	result := add(3, 4)
	fmt.Println("result ", result)
	// ====================
	multiplication := mul(3, 4)
	fmt.Println("multiplication ", multiplication)

	// ================================
	value1, value2, value3, value4, value5 := getMultipleData()
	fmt.Println(value1, value2, value3, value4, value5)

	// ================================

	fn := func(v int) int {
		return 9
	}

	processIt(fn)

}

