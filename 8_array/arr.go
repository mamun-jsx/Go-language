package main

import "fmt"

// numbered  sequence of specific length
func main() {
	var numbs [4]int
	// add something into array  into 0 index
	numbs[0] = 100
	// print the 1st index value
	fmt.Println("arr 1st idx value", numbs[0])
	// inbuilt feature to check length
	fmt.Println(len(numbs))
	// print full array

	fmt.Println("full arr", numbs)
	// when you declare an arr with specific length
	// and golang will create all 00000 for you

	// ? declare arr and also initial the value
	//? declare into single line
	nums := [3]int{1, 3, 5}
	fmt.Println("instant define : ", nums)
	//*  2d arr

	numberss := [2][3]int{{1, 3, 5}, {2, 4, 6}}
	fmt.Println("2d arr", numberss)

}
