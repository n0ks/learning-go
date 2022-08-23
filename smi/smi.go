package smi

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
  Base float64
  Height float64
  
}
func Perimeter(rect Rectangle) float64 {

	return 2 * (rect.width + rect.height)

}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) foo() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
