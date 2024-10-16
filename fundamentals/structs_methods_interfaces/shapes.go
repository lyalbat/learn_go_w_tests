package main

import "math"

// parametric polymorphism
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// syntax for declaring methods
func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
