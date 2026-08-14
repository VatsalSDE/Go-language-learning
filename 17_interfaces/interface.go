package main

import "fmt"

// type paymenter interface {
// 	pay(amount float32)
// 	refund(amount float32, account string)
// }

// type payment struct {
// 	gateway         stripe
// 	razorpaygateway razorpay
// }

// // Open close principle is violated here !!
// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGw := razorpay{}
// 	// stripePaymentGw := stripe{}
// 	// razorpayPaymentGw.pay(amount)
// 	// stripePaymentGw.pay(amount)
// 	p.gateway.pay(amount)
// 	p.razorpaygateway.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	// logic to make payment
// 	fmt.Println("making payment using razorpay", amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

// type fakepayment struct{}

// func (f fakepayment) pay(amount float32) {
// 	fmt.Println("making payment using fake gateway for testing purpose")
// }

// type paypal struct{}

// func (p paypal) pay(amount float32) {
// 	fmt.Println("making payment using paypal", amount)
// }

// func (p paypal) refund(amount float32, account string) {

// }

// func main() {
// 	// stripePaymentGw := stripe{}
// 	// razorpayPaymentGw := razorpay{}

// 	// fakeGw := fakepayment{}
// 	// paypalGw := paypal{}
// 	// newPayment := payment{
// 	// 	gateway: paypalGw,
// 	// }

// 	newPayment := payment{}
// 	newPayment.makePayment(100)
// }

// this is the contract

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// this is breaking the open close principle as we are adding new payment gateway in the existing code
func (p payment) makePayment(amount float32) {
	// razorpayinstance := razoray{}
	// razorpayinstance.pay(amount)
	// stripeinstance := stripe{}
	// stripeinstance.pay(amount)

	// nowusing the instance we have created likewise
	p.gateway.pay(amount)
	// p.stripegateway.pay(amount) // but this is also breaking the open close principle as we are adding new payment gateway in the existing code
	// which is still breaking the open close principle as we are adding new payment gateway in the existing code
	// so we need to create a new interface and implement it in the payment struct

}

// now making for the razoray the struct
type razoray struct {

}

func (r razoray) pay(amount float32) {

	fmt.Println("Razorpay is doing the payment of ", amount)

}


type stripe struct {
}

func (s stripe) pay(amount float32) {

	fmt.Println("Stripe is doing the payment of ", amount)

}

func main() {

	// stripepayment := stripe{}
	razorpaypayment := razoray{}
	newPayment := payment{
		gateway: razorpaypayment,
	}

	newPayment.makePayment(100)

}
