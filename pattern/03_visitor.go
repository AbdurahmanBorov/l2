package pattern

import "fmt"

type Shape interface {
	Accept(Visitor)
}

type Circle struct{}

func (c *Circle) Accept(v Visitor) { v.VisitCircle(c) }

type Square struct{}

func (s *Square) Accept(v Visitor) { v.VisitSquare(s) }

type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	fmt.Println("Рассчитываем площадь круга")
}
func (a *AreaCalculator) VisitSquare(s *Square) {
	fmt.Println("Рассчитываем площадь квадрата")
}

func VisitorPatter() {
	shapes := []Shape{&Circle{}, &Square{}}
	calculator := &AreaCalculator{}

	for _, shape := range shapes {
		shape.Accept(calculator)
	}
}
