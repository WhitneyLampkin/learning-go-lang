package main

/*
	fmt
	math
	os - command-line arguments
	strconv - string conversations
*/

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/whitneylampkin/learning-go-lang/calculator"
)

func main() {
	fmt.Println("Hello World!")

	// Referencing a local package
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

	mostCommonVariableDeclaration()
	declaringVariables()
	initializingVariables()
	integerDataType()
	floatDataType()
	stringDataType()
	defaultValues()
	typeConversations()
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

/*
	Variables: The following code will throw an error because variables are declared but not used.

	func main() {
		firstName, lastName := "John", "Doe"
		age := 32
	}

	./main.go:4:2: firstName declared but not used
	./main.go:4:13: lastName declared but not used
	./main.go:5:2: age declared but not used
*/

func mostCommonVariableDeclaration() {
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)
}

func declaringVariables() {
	// Declaring a single variable
	var myFirstName string
	myFirstName = "Whitney"
	fmt.Println(myFirstName)

	// Declaring multiple variables
	var hisFirstName, hisLastName string
	hisFirstName = "Regis"
	hisLastName = "Davis Jr."
	fmt.Println(hisFirstName, hisLastName)

	// Declaring multiple variable with different data types
	var herFirstName, herLastName string
	var herAge int
	herFirstName = "Kai"
	herLastName = "Lymore"
	herAge = 9
	fmt.Println(herFirstName, herLastName, herAge)

	// Simplified variable declaration
	var (
		firstName, lastName string
		age                 int
	)
	firstName = "First"
	lastName = "Last"
	age = 35
	fmt.Println(firstName, lastName, age)
}

func initializingVariables() {
	// Initializing variables with data types
	var (
		firstNameJohn string = "John"
		lastNameDoe   string = "Doe"
		age32         int    = 32
	)
	fmt.Println(firstNameJohn, lastNameDoe, age32)

	// Initializing variables without data types
	var (
		firstNameJohn2 = "John"
		lastNameDoe2   = "Doe"
		age33          = 33
	)
	println(firstNameJohn2, lastNameDoe2, age33)

	// Declaring and initializing variables in a single line
	var (
		firstNameJohn3, lastNameDoe3, age34 = "John", "Doe", 34
	)
	println(firstNameJohn3, lastNameDoe3, age34)
}

/*
	Data Types: Go is a strongly typed language so variables will only take values with the correct data type.

	- Basic types: numbers, strings, and booleans
	- Aggregate types: arrays and structs
	- Reference types: pointers, slices, maps, functions, and channels
	- Interface types: interface
*/

func integerDataType() {
	// Various integer types
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	/*
		- The following code throws an error because the two types don't match.
		- Must explicitly define the new type.

		var integer16 int16 = 127
			var integer32 int32 = 32767
			fmt.Println(integer16 + integer32)
		}

		Throws the following error: invalid operation: integer16 + integer32 (mismatched types int16 and int32)
	*/

	// Runes: "A rune is simply an alias for int32 data type. It's used to represent a Unicode character (or a Unicode code point)."
	// The following rune prints 71, the unicode character for G.
	rune := 'G'
	fmt.Println(rune)

	// Challenge 1
	var integer32Again int = 2147483648
	var integerVariable = integer64
	fmt.Println(integer32Again, integerVariable)

	/*
		Challenge 2: Produces an error "const -10 overflows uint"
			var integer uint = -10
			fmt.Println(integer)
	*/
}

func floatDataType() {
	// Floating-point numbers can be used for values larger than ints and decimals
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	fmt.Println(e, Avogadro, Planck)
}

func booleanDataType() {
	var featureFlag bool = true
	fmt.Println(featureFlag)
}

func stringDataType() {
	/*
		"" - Used for strings
		'' - Used for single characters and runes
	*/

	var firstName string = "John"
	lastName := "Doe"
	fmt.Println(firstName, lastName)

	/*
		Escape Characters

		\n for new lines
		\r for carriage returns
		\t for tabs
		\' for single quotation marks
		\" for double quotation marks
		\\ for backslashes
	*/

	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)
}

func defaultValues() {
	/*
		0 for int types (and all of its subtypes, like int64)
		+0.000000e+000 for float32 and float64 types
		false for bool types
		An empty value for string types
	*/

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
}

func typeConversations() {
	// Type conversations using casting
	var integer16Bit int16 = 127
	var integer32Bit int32 = 32767
	fmt.Println(int32(integer16Bit) + integer32Bit)

	// Type conversations using strconv()
	// Take note of the _
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
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
