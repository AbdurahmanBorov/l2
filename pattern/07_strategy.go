package pattern

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Оплата %.2f с помощью кредитной карты\n", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Оплата %.2f через PayPal\n", amount)
}

type ShoppingCart struct {
	strategy PaymentStrategy
}

func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	s.strategy = strategy
}

func (s *ShoppingCart) Checkout(amount float64) {
	s.strategy.Pay(amount)
}

func Strategy() {
	cart := &ShoppingCart{}

	cart.SetPaymentStrategy(&CreditCardPayment{})
	cart.Checkout(100)

	cart.SetPaymentStrategy(&PayPalPayment{})
	cart.Checkout(200)
}
