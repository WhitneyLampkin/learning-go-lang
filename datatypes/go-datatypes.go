package datatypes

import (
	"fmt"
	"math"
	"strconv"
)

/*
	Data Types: Go is a strongly typed language so variables will only take values with the correct data type.

	- Basic types: numbers, strings, and booleans
	- Aggregate types: arrays and structs
	- Reference types: pointers, slices, maps, functions, and channels
	- Interface types: interface
*/

func IntegerDataType() {
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

func FloatDataType() {
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

func BooleanDataType() {
	var featureFlag bool = true
	fmt.Println(featureFlag)
}

func StringDataType() {
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

func DefaultValues() {
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

func TypeConversations() {
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
