package main

import "fmt"

func sum(numbs ...int) int {
	total := 0
	for _, num := range numbs {
		total = total + num
	}
	return total

}

func main() {
	result := sum(1, 2, 3)

	fmt.Println("result ", result)
	// make an slice
	numberss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("result ", sum(numberss...))

}
