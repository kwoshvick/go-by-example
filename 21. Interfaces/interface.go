package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

type rectangle struct {
	length, width float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * r.width * r.length
}

func main() {

	gC := geometry(circle{7})
	fmt.Println(gC.area())
	fmt.Println(gC.perimeter())

	gR := geometry(rectangle{5, 12})
	fmt.Println(gR.area())
	fmt.Println(gR.perimeter())

}
