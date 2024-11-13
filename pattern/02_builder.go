package pattern

import "fmt"

type Car struct {
	Wheels int
	Color  string
	Model  string
}

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{&Car{}}
}

func (b *CarBuilder) SetWheels(wheels int) *CarBuilder {
	b.car.Wheels = wheels
	return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) SetModel(model string) *CarBuilder {
	b.car.Model = model
	return b
}

func (b *CarBuilder) Build() *Car {
	return b.car
}

func Builder() {
	car := NewCarBuilder().
		SetWheels(4).
		SetColor("Красный").
		SetModel("Спорт").
		Build()
	fmt.Printf("Машина: %+v\n", car)
}
