package main

import "fmt"

// Use a struct to represent a rectangle
type Rectangle struct {
	length float64
	width  float64
}

// These are called methods
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	// First rectangle
	rec1 := Rectangle{
		length: 4.5,
		width:  2,
	}

	// Second rectangle
	rec2 := Rectangle{
		length: 4.5,
		width:  2.5,
	}

	fmt.Println(rec1.Area())
	fmt.Println(rec2.Perimeter())

}
