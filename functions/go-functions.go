package functions

import (
	"fmt"
	"os"
	"strconv"
)

func MainFunctions() {
	// Summing numbers from cmd line
	// Atoi = parseInt
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum:", number1+number2)
}

func CustomSumFunction(number1 string, number2 string) int {
	/*
		func name(parameters) (results) {
			body-content
		}
	*/

	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

func ReturnMultipleValues(number3 string, number4 string) (sum int, mul int) {
	int3, _ := strconv.Atoi(number3)
	int4, _ := strconv.Atoi(number4)
	sum = int3 + int4
	mul = int3 * int4
	return
}

func PointerParametervalues() {
	// Go is a "pass by value" programming language.
	// Functions make their own local copies of args that do not change the original value.
	// The updateName() function changes the variable's value locally but this function will still print "John".
	firstName := "John"
	updateName(firstName)
	fmt.Println(firstName)

	// Pointers are used to change the original value
	// "A pointer is a variable that contains the memory address of another variable."
	updateNamePointer(&firstName)
	fmt.Println(firstName)
}

func updateName(name string) {
	name = "David"
}

func updateNamePointer(name *string) {
	*name = "Paul"
}
