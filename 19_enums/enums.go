package main

import "fmt"

// enum types is for declare multiple types. into go lang
// there is no enum like typescript into golang we use const and declare multiple

type OrderStatus int

const (
	Recived OrderStatus = iota
	Confirmed
	Shipped
	Delivered
	Canceled
)

type DeleteStatus string

const (
	Yes DeleteStatus = "yes"
	No  DeleteStatus = "no"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("change order status to ", status)
}
func deleteOrderStatus(status DeleteStatus) {
	fmt.Println("delete status ", status)
}
func main() {

	changeOrderStatus(Delivered)
	deleteOrderStatus(Yes)
}
