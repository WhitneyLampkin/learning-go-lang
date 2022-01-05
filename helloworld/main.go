package main

/*
	fmt
	math
	os - command-line arguments
	strconv - string conversations
*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/whitneylampkin/learning-go-lang/calculator"
	"github.com/whitneylampkin/learning-go-lang/datatypes"
	"github.com/whitneylampkin/learning-go-lang/variables"
)

func main() {
	fmt.Println("Hello World!")

	// Referencing a local package
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

	variables.MostCommonVariableDeclaration()
	variables.DeclaringVariables()
	variables.InitializingVariables()
	datatypes.IntegerDataType()
	datatypes.FloatDataType()
	datatypes.StringDataType()
	datatypes.DefaultValues()
	datatypes.TypeConversations()
	mainFunctions()

	// Calling customSumFunction()
	sum := customSumFunction(os.Args[1], os.Args[2])
	println("Sum with a custom function:", sum)

	// Calling returnMultipleValues()
	sum, mul := returnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	// Discarding variables
	sum2, _ := returnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum2 (Discarding mul):", sum2)

	pointerParametervalues()
}

func mainFunctions() {
	// Summing numbers from cmd line
	// Atoi = parseInt
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum:", number1+number2)
}

func customSumFunction(number1 string, number2 string) int {
	/*
		func name(parameters) (results) {
			body-content
		}
	*/

	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

func returnMultipleValues(number3 string, number4 string) (sum int, mul int) {
	int3, _ := strconv.Atoi(number3)
	int4, _ := strconv.Atoi(number4)
	sum = int3 + int4
	mul = int3 * int4
	return
}

func pointerParametervalues() {
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
