package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	// customer
}

// this is the constructor

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

// func (o order) getAmount() float32 {
// 	ret urn o.amount
// }

func main() {
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	// newOrder := order{
	// 	id:     "1",
	// 	amount: 30,
	// 	status: "received",
	// 	customer: customer{
	// 		name:  "john",
	// 		phone: "1234567890",
	// 	},
	// }

	// newOrder.customer.name = "robin"
	// fmt.Println(newOrder)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// myOrder2 := myOrder

	// myOrder2.amount = 20

	// fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder2.status = "paid"

	fmt.Println("Order struct", myOrder2.id, myOrder2.amount, myOrder2.status, myOrder2.createdAt)
	// fmt.Println("Order struct", myOrder)
}
