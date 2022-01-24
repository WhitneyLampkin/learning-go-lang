package main

import (
	"fmt"
	/*

		Uncomment to use all of the functions in main()

		"os"

		"github.com/whitneylampkin/learning-go-lang/calculator"
		"github.com/whitneylampkin/learning-go-lang/controlflows"
		"github.com/whitneylampkin/learning-go-lang/datatypes"
		"github.com/whitneylampkin/learning-go-lang/forloops"
		"github.com/whitneylampkin/learning-go-lang/functions"
		"github.com/whitneylampkin/learning-go-lang/variables" */)

func main() {
	/* fmt.Println("Hello World!")

	// Referencing a local package
	fmt.Println("***Referencing a local package***")
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

	// Referencing 3rd-party packages
	fmt.Println("***Referencing 3rd party packages***")
	//fmt.Println(quote.Hello())

	// Variables, datatypes and functions
	fmt.Println("***Variables, datatypes and functions***")
	variables.MostCommonVariableDeclaration()
	variables.DeclaringVariables()
	variables.InitializingVariables()
	datatypes.IntegerDataType()
	datatypes.FloatDataType()
	datatypes.StringDataType()
	datatypes.DefaultValues()
	datatypes.TypeConversations()
	functions.MainFunctions()

	// Calling customSumFunction()
	fmt.Println("***Calling custom functions***")
	sum := functions.CustomSumFunction(os.Args[1], os.Args[2])
	println("Sum with a custom function:", sum)

	// Calling returnMultipleValues()
	fmt.Println("***Returning multiple values***")
	sum, mul := functions.ReturnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	// Discarding variables
	fmt.Println("***Discarding variables***")
	sum2, _ := functions.ReturnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum2 (Discarding mul):", sum2)

	functions.PointerParametervalues()

	// Control Flow
	fmt.Println("***Control Flow***")
	controlflows.IfStatements()
	controlflows.CompoundIfStatements()
	controlflows.SwitchStatements()
	controlflows.SwitchMultipleExpressions()
	controlflows.SwitchInvokeFunction()
	controlflows.SwitchRegEx()
	controlflows.OmitConditions()
	controlflows.SwitchBreaks()
	controlflows.DeferExample()
	// The next line is commented out because it's stops execution.
	// controlflows.PanicHighLowExample()
	controlflows.RecoverExample()

	// Forloops
	forloops.BasicForLoop()
	forloops.FakeWhileLoop()
	forloops.InfiniteLoop()

	// Challenge: FizzBuzz
	Fizzbuzz()

	// Challenge: Find the primes (under 20)
	fmt.Println("Prime numbers less than 20:")

	for number := 2; number < 20; number++ {
		if findprimes(number) == true {
			fmt.Printf("%v ", number)
		}
	}

	// Challenge: Ask for a number, panic if negative
	val := 0
	fmt.Print("\r\nEnter number: ")
	fmt.Scanf("%d", &val)

	switch {
	case val < 0:
		panic("The program is crashing because you entered a negative number.")
	case val > 0:
		fmt.Println("You entered:", val)
	case val == 0:
		fmt.Println("0 is neither negative nor positive")

	} */

	// Arrays
	DeclareArrays()
	InitializingArrays()
	EllipsisInArrays()
	ArrayLastPositionOnly()
	MultidimensionalArrays()
}

func Fizzbuzz() {
	for i := 1; i <= 100; i++ {
		// TODO: Use switch case instead
		// Could also use i%15 for the first case
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
			fmt.Println(" - FizzBuzz")
		}
		if i%3 == 0 && i%5 != 0 {
			fmt.Println(i)
			fmt.Println(" - Fizz")
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println(i)
			fmt.Println(" - Buzz")
		}
	}

}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

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
