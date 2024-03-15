package main

import (
	"fmt"
)

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{
		Width:  2,
		Height: 3,
	}
	t := Triangle{Base: 4, Height: 6}

	fmt.Println("Original Shape Circle:")
	PrintShapeProperties(c)
	fmt.Println("Original Shape Rectangle:")
	PrintShapeProperties(r)
	fmt.Println("Original Shape Triangle:")
	PrintShapeProperties(t)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.1416 * c.Radius * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.1416 * c.Radius
}

func (t Triangle) Perimeter() float64 {
	return t.Base * 4
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func PrintShapeProperties(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}
