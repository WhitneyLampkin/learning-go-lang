package arrays

import "fmt"

func DeclareArrays() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
}

func InitializingArrays() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}

func EllipsisInArrays() {
	cities2 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities 2:", cities2)
}

// Initializing an array in this way will create an array with a length of 100
// Only the last position will have a value set. The others will default to 0, which is the default for type int.
func ArrayLastPositionOnly() {
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}

func MultidimensionalArrays() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}
