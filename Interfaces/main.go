package main

import (
	"fmt"
	"math"
)

// Interfaces are basically contracts for instance of methods of a specific struct
// this provides a high level organisation without really implementing the business logic
// Structs should implement these methods

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Rectangle Satisfies Measurable and SHape interfaces both
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Rectangle Satisfies Measurable interface as well
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter", g.Perimeter())
}

type Person struct {
	Id   string
	Name string
}

type PaymentProcessor interface {
	CheckoutPayment()
}

type StripePaymentProcessor struct {
	URL string
}

type MockPaymentProcessor struct {
	URL string
}

func (p StripePaymentProcessor) CheckoutPayment() {
	fmt.Println("Stripe Payment Processing", p.URL)
}

func (p MockPaymentProcessor) CheckoutPayment() {
	fmt.Println("Mock Payment Processing", p.URL)
}

func Checkout(p PaymentProcessor) {
	p.CheckoutPayment()
}

func main() {
	fmt.Println("HELLO FROM INTERFACES")

	r := Rectangle{width: 5, height: 10}

	// We are able to pass Rectangle because it implicitly satisfies the Geometry interface.
	// Geometry is composed of Shape and Measurable interfaces,
	// which together require Area() and Perimeter() methods.
	// Rectangle implements both methods, so it satisfies Geometry.

	describeShape(r)

	data := map[string]any{
		"id":   1,
		"name": "Priyanshu",
		"person": Person{
			Id:   "1",
			Name: "Priyanshu Naskar",
		},
	}

	for key, val := range data {
		switch v := val.(type) {
		case int:
			fmt.Printf("%s is an int: %d\n", key, v)
		case string:
			fmt.Printf("%s is a string: %s\n", key, v)
		case Person:
			fmt.Printf("%s is a Person: Name=%s\n", key, v.Name)
		default:
			fmt.Printf("%s has unknown type %T\n", key, v)
		}
	}

	dev := "DEV"
	if dev == "DEV" {
		mock_payment_processor := MockPaymentProcessor{
			URL: "MOCK PAYMENT LINK",
		}

		Checkout(mock_payment_processor)

	} else {
		stripe_payment_processor := StripePaymentProcessor{
			URL: "STRIPE PAYMENT LINK",
		}

		Checkout(stripe_payment_processor)
	}

}
