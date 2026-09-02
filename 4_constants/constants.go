package main

import "fmt"

func main() {
	// unchangeable  variable we can declare it outside of main func
	//  but  if you use short hand name:="my name" it will not work
	const name string = "Go Lang"
	fmt.Println(name)
	// ! multiple constant
	const (
		userName = "School-management"
		host     = "localhost"
		port     = 4001
	)
	fmt.Println(userName,port,host)
}
