package controlflows

import (
	"fmt"
	"math/rand"
	"time"
)

// If statements do not need ().
// Ternary if statements are not supported.
func IfStatements() {
	x := 27
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	if x%2 != 0 {
		fmt.Println(x, "is odd")
	}
}

/*
	This following code produces an error:

	if num := somenumber(); num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    fmt.Println(num)
*/

func CompoundIfStatements() {
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func givemeanumber() int {
	return -1
}

// If you run the following code several times, you'll see a different output every time. (But if you run the code in the Go Playground, you'll get the same result every time. That's one of the service's limitations.
func SwitchStatements() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	}

	fmt.Println("ok")
}

func SwitchMultipleExpressions() {

}

func SwitchInvokeFunction() {

}

func OmitConditions() {

}

func SwitchBreaks() {

}
