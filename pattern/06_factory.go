package pattern

import "fmt"

type Transport interface {
	Deliver()
}

type Truck struct{}

func (t *Truck) Deliver() { fmt.Println("Доставка грузовиком") }

type Ship struct{}

func (s *Ship) Deliver() { fmt.Println("Доставка кораблём") }

type LogisticsFactory interface {
	CreateTransport() Transport
}

type RoadLogistics struct{}

func (r *RoadLogistics) CreateTransport() Transport { return &Truck{} }

type SeaLogistics struct{}

func (s *SeaLogistics) CreateTransport() Transport { return &Ship{} }

func Factory() {
	var logistics LogisticsFactory
	logistics = &RoadLogistics{}
	logistics.CreateTransport().Deliver()
}
