package main

// func main() {
// 	_ = Shape.Area()
// 	return area
// }

// Shape is an interface that requires an Area method returning a float64.
// Any type that has an Area() float64 method satisfies this interface.
type Shape interface {
	Area() float64
}

// Rectangle is a struct with width and height fields.
type Rectangle struct {
	width, height float64
}

// Area is a method on *Rectangle that calculates and returns the area.
// The receiver is a pointer, allowing interface implementation and avoiding value copying.
func (r *Rectangle) Area() float64 {
	return r.width * r.height // Correct area formula: width × height
}
