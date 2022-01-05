package controlflows

import "fmt"

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

func SwitchStatements() {

}

func SwitchMultipleExpressions() {

}

func SwitchInvokeFunction() {

}

func OmitConditions() {

}

func SwitchBreaks() {

}
