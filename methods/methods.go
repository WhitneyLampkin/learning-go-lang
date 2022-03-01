package methods

import "fmt"

func DeclareMethod() {
	t := triangle{3}
	s := square{4}
	fmt.Println("\r\nPerimeter (triangle):", t.perimeter())
	fmt.Println("\r\nPerimeter (square):", s.perimeter())
}

// Structs
type triangle struct {
	size int
}

type square struct {
	size int
}

// Methods

// This method requires a receiver, the triangle struct.
func (t triangle) perimeter() int {
	return t.size * 3
}

// Go allows methods with the same name as long as the receivers are different.
func (s square) perimeter() int {
	return s.size * 4
}
