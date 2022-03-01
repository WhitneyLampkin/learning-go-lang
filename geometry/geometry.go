package geometry

// This is package is an encapsulation example from the Methods lesson.

// Public
type Triangle struct {
	size int
}

// Private
func (t *Triangle) doubleSize() {
	t.size *= 2
}

// Public
func (t *Triangle) SetSize(size int) {
	t.size = size
}

// Public
func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}
