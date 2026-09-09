package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}
func job(ids int){
	fmt.Println("doing job ", ids)
}

func main() {
	for i := 1; i <= 10; i++ {
		go task(i)
	}

	for i:=30; i<=40 ; i++{
		go func(ids int){
			fmt.Println("doing job",ids)
		}(i)
	}

	// do another way or syntex 

	time.Sleep(time.Second * 2)
}
