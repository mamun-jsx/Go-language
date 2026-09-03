package main

import "fmt"

// iterating over data structure
// ? Go's range and JavaScript's for...of are highly similar
// ? because both are used to iterate over elements in a collection without managing a traditional loop counter.
func main() {
	nums := []int{6, 7, 8}

	// fmt.Println(nums)
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// * use range with for
	sum := 0
	for index, num := range nums {
		sum = num + 1
		fmt.Println(sum, index)
	}
	m := map[string]string{"name": "john", "lName": "Doe"}

	// print the m-> map by for loop with range
	for key, value := range m {
		fmt.Println(key, value)
	}
	// range use into string
	// * i see binary code of each char
	// *
	for i, c := range "golang" {
		// fmt.Println(i, c)
		// show the char
		fmt.Println(i, string(c))
		// when use string(c) the code is shown to human readable
	}
}
