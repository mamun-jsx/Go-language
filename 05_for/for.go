package main

import "fmt"

// for --> only construct into go for looping
func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// this is an infinity loop
	// for {
	// 	fmt.Println("run ")
	// }

	//? classic for loop
	for i := 0; i <= 10; i++ {
		//skip any number
		if i == 2 {
			continue
		}
		fmt.Println("show", i)
	}
	// range
	fmt.Printf("_____range________")
	// range is similar like for of loop into js which reads into 1st to last each element
	for i := range 5 {
		fmt.Println(i)
	}
	// ? ternary operator : I have to use normal if else 

}
