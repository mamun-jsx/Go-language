package main

import (
	"fmt"
	"time"
)

// Structs is similar like class
// in struct we can define variables of different types
// make complex data structure

// order structs for an ecommerce

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // it comes from another packedge and working with nenosecond precision
}

// need to make an instance -> my order and myOrder2

// connect with struct
func (o *order) changeStatus(status string) {
	// it is reviver type
	o.status = status
	// we need to pass the memory location via pointer
}

func (o order) getAmount() float32 {
	return o.amount
	// to get anything from struct we do not need to dereference it
	// to change it we have to send the pointer as memory location
	// if you want to modify anything you need to pass memory location

}

// constractor into golang as opp code
func NewOrder(id string, amount float32, status string) *order {
	myOrder3 := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder3
}

func main() {
	// Anonymous struct
	person := struct {
		name   string
		isGood bool
	}{name: "goLang", isGood: true}

	myOrder := order{
		id:     "10",
		amount: 43,
		status: "received",
	}
	myOrder.changeStatus("Confirmed") // reciver mathods to chance any value
	// into struct we do not need to dereference because struct do it automaticallly
	fmt.Println("change status reciver", myOrder)
	fmt.Println(myOrder.status)
	myOrder.createdAt = time.Now()
	fmt.Println("Order struc ", myOrder)
	fmt.Println("getAmount ", myOrder.getAmount())
	// ================ make another instance===========
	myOrder2 := order{
		id:        "11",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
	}
	myOrder.status = "shipped" // value will change into my order and do not effect into myOrder2 instance
	fmt.Println("myOrder2 ", myOrder2)

	// ====== constractor
	myNewOrder := NewOrder("100", 30, "Not Paid")
	fmt.Println("my newOrder ", myNewOrder)
	// using anonymous struct
	fmt.Println("language anonymous structs ", person)
}

// struct is working as a blue print of a dataStructure
