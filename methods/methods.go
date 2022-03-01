package methods

import "fmt"

func DeclareMethod() {
	t := triangle{3}
	fmt.Println("Perimeter:", t.perimeter())
}

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}
