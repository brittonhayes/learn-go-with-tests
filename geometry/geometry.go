package geometry

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	Base       float64
	Height     float64
	Hypotenuse float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius   float64
	Diameter float64
}

// Perimeter returns the perimeter provided the width and
// height of an object
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a rectangle given the
// width and height
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area of a triangle given
// the base and height
func (t Triangle) Area() float64 {
	return t.Base * t.Height
}

// Perimeter returns the perimeter of a triangle
// given the base, height, and hypotenuse
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + t.Hypotenuse
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of
// a circle give the diameter
func (c Circle) Perimeter() float64 {
	return math.Pi * c.Diameter
}
