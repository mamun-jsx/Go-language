package main

import "fmt"

func main() {
	// check condition
	age := 17
	if age >= 18 {
		fmt.Println("person can enter into room")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person permission deny")
	}
	//  logical operator permission
	var role string = "admin"
	var hasPermission bool = false
	if role == "admin" && hasPermission {
		fmt.Println("person has permission")
	} else {
		fmt.Printf("Bad auth permission deny")
	}
}
