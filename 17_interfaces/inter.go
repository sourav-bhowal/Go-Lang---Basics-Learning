package main

import "fmt"

// Paymentr is an interface that defines a method for processing payments.
type Paymentr interface {
	Pay(amount float32) // Pay method to process a payment with a given amount
}

// Razorpay is a struct that implements the Pay method for payment processing.
type Razorpay struct{}

// Pay is a method of Razorpay that processes a payment.
func (r Razorpay) Pay(amount float32) {
	fmt.Printf("Payment of %.2f made through Razorpay\n", amount)
}

// Stripe is a struct that implements the Pay method for payment processing.
type Stripe struct{}

// Pay is a method of Stripe that processes a payment. It is coming from the Paymentr interface. Go on its own implements the interface for us. Just need to define the method as required by the interface.
func (s Stripe) Pay(amount float32) {
	fmt.Printf("Payment of %.2f made through Stripe\n", amount)
}

// Payment is a struct that contains a Paymentr interface, allowing it to use any payment gateway that implements the Pay method.
type Payment struct {
	gateway Paymentr // Paymentr is an interface that defines the Pay method for payment processing
}

// Pay is a method of Payment that calls the Pay method of Razorpay.
func (p Payment) MakePayment(amount float32) {
	p.gateway.Pay(amount)
}

func main() {

	// Stripe Gateway
	stripeGateway := Stripe{}
	// Razorpay Gateway
	// razorpayGateway := Razorpay{}

	// Create an instance of Payment
	myPayment := Payment{
		gateway: stripeGateway, // Assign the Stripe gateway to the Payment instance
		// gateway: razorpayGateway, // Assign the Razorpay gateway to the Payment instance
	}

	// Call the MakePayment method to process a payment
	myPayment.MakePayment(100.50) // Example amount to be paid
}
