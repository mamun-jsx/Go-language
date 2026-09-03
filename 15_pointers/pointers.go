package main

import "fmt"

// pointer -> very important for low level memory management
// pointer -> it store the address of variable

func changeNum(num int) {
	num = 5
	fmt.Println("in changeNum", num)
}

// if we need to change the value we have to send the referance | it means I have to send the memory location
// pointers does the same thing

// ======== Pointer========
func changAbleNumber(numbe *int) {
	// send the number's referance as the memory location
	// pointer symbol * (star)
	*numbe = 15 // adding here * is dereferencing (means change the value)

	fmt.Println("changable number ", numbe)
}

func main() {

	num := 1
	changeNum(num)
	fmt.Println("memory location ", &num)
	// output : memory location  0x1a7b01e720b0
	// this is the memory location
	// fmt.Println("after change num into main ", num)

	toChangeNumbe := 100
	fmt.Println("before change toChangeNumbe into main ", toChangeNumbe)

	changAbleNumber(&toChangeNumbe)
	fmt.Println("after change toChangeNumbe into main ", toChangeNumbe)
}
