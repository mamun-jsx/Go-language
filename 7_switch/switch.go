package main

import (
	"fmt"
	"time"
)

// when we use switch statement
//
//	when logic is very complex
func main() {
	i := 0
	// into switch , you do not need to use break,
	// go by default handle the break
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other number")
	}
	// multiple condition switch
	switch time.Now().Weekday() {
	// time comes from package like fmt
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("Working day")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("value an integer", v)
		case string:
			fmt.Println("value is string")
		case bool:
			fmt.Println("value is boolean ")
		default:
			fmt.Println("something else")
		}
	}
	whoAmI(nil)
}
