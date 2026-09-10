package main

import (
	"fmt"
	"slices"
)

// it is an dynamic array
// do not need to use fixed length
// + useful methods
func main() {
	// declare a slice
	var num []int
	fmt.Println("see the num", num) //show empty as nil
	fmt.Println(len(num))           // length 0 because no value is assigned

	var nums = make([]int, 2, 3) //make(...): The built-in memory allocation tool.
	// capacity check
	// fmt.Println("maximum numbers of elements can fit", cap(nums))
	// fmt.Println(nums)
	nums = append(nums, 1) // add last into the index
	nums = append(nums, 2) // add last into the index
	nums = append(nums, 3) // add last into the index
	nums = append(nums, 4) // add last into the index
	nums = append(nums, 5) // add last into the index
	fmt.Println(nums)
	fmt.Println("capacity", cap(nums))

	/*

		! How capacity grows here:
		1. Starts with len=2, cap=3 [0, 0]
		2. Adding 1 & 2 fits perfectly -> cap stays 3
		3. Adding 3 breaks the limit -> Go doubles cap to 6 behind the scenes
		4. Adding 4 & 5 fits into the new cap of 6
	*/

	// another way to define the slice
	numberss := []int{}
	numberss = append(numberss, 1)
	numberss = append(numberss, 2)
	fmt.Println(cap(numberss), "check capacity of this slice")
	fmt.Println(numberss, "see full slice")

	// *slice operator
	var numb = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(numb[4:7], "checking index number") // 1st from index 0 to last input the number of index to stop as return

	// slice
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2), "check is they are Equal or not")
	// slices is also a inbuilt system like fmt

}
