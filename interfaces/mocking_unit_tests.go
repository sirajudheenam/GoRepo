package main

import "fmt"

// Interfaces make it easy to mock dependencies.
// Useful for separating business logic from external systems.
type PaymentProcessor interface {
	Pay(amount float64) bool
}

type RealPayment struct{}

func (RealPayment) Pay(amount float64) bool { return true }

type FakePayment struct{}

func (FakePayment) Pay(amount float64) bool { return false }

func Checkout(p PaymentProcessor) {
	success := p.Pay(99.99)
	if success {
		fmt.Println("Payment successful")
	} else {
		fmt.Println("Payment failed")
	}
}

func main() {
	Checkout(RealPayment{})
	Checkout(FakePayment{})
}
