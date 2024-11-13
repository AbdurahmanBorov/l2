package pattern

import "fmt"

type PaymentService struct{}

func (p *PaymentService) Pay(amount float64) { fmt.Println("Оплата на сумму:", amount) }

type ShippingService struct{}

func (s *ShippingService) Ship(item string) { fmt.Println("Доставка товара:", item) }

type OrderFacade struct {
	payment  *PaymentService
	shipping *ShippingService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		payment:  &PaymentService{},
		shipping: &ShippingService{},
	}
}

func (f *OrderFacade) PlaceOrder(amount float64, item string) {
	f.payment.Pay(amount)
	f.shipping.Ship(item)
}

func Facade() {
	order := NewOrderFacade()
	order.PlaceOrder(100, "Книга")
}
