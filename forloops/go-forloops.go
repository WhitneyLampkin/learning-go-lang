package forloops

import "fmt"

// Basic for loop syntax
func BasicForLoop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum of 1..100 is", sum)
}
