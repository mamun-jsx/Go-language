package main

import "fmt"

func printSlice[T int | string](item []T) {
	for _, v := range item {
		fmt.Println(v)
	}
}
func printSlices[T comparable](item []T) {
	for _, v := range item {
		fmt.Println(v)
	}
}

type stack[T string | int] struct {
	element []T
}

func main() {

	myStack := stack[int]{
		element: []int{1, 2, 3, 4},
	}
	myLange := stack[string]{
		element: []string{"go lang", "js", "c++"},
	}

	fmt.Println(myStack.element)
	fmt.Println(myLange.element)

	name := []string{"go lang", "js", "c++"}
	
	// Using the original generic function
	printSlice([]int{1, 2, 3}) 
	printSlice(name)           

	// Using the comparable generic function
	fmt.Println("--- Using comparable ---")
	printSlices([]float64{3.14, 2.71})
}
