package main

import (
	"fmt"
	"maps"
)

// map--> a data structure like hash, object something like this
func main() {
	// =========================================================================
	// GO MAP QUICK CHEAT SHEET: make(map[KeyType]ValueType)
	// Inside [ ]: The Lookup Key (e.g., string, int, bool)
	// Right side: The Actual Value stored (e.g., string, int, struct)
	// =========================================================================
	// creating map
	m := make(map[string]string)
	// setting an element into this m
	m["name"] = "Go Lang"
	m["area"] = "backend"
	// get an element
	fmt.Println(m["name"], m["area"])
	// fmt.Println(m["photos"])
	n := make(map[string]int)
	n["age"] = 30
	n["feature"] = 3
	n["price"] = 3000

	delete(n, "price") // delete a single item
	clear(n)           // it will clear the map
	fmt.Println(n)

	fmt.Println(len(n), "check the length of the n map")

	//============ create map without make()===========

	withoutMake := map[string]int{"price": 40, "phone": 4}
	fmt.Println(withoutMake, "without map create")
	// The 'ok' variable is a boolean (true/false) that explicitly tells you
	// if the key exists in the map, preventing confusion with default 0 values.
	value, ok := withoutMake["phone"]
	fmt.Println("value of withoutMake", value)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// ============|Check Equal or NOt|======================
	l := map[string]int{"price": 1, "phone": 3}
	g := map[string]int{"price": 1, "phone": 2}
	fmt.Println(maps.Equal(l, g)) // return false because they are not same.
	// map is the data-structure BUT //? maps is the package like fmt
}
