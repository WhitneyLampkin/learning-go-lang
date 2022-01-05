package main

import (
	"fmt"
	"os"

	"rsc.io/quote"

	"github.com/whitneylampkin/learning-go-lang/calculator"
	"github.com/whitneylampkin/learning-go-lang/datatypes"
	"github.com/whitneylampkin/learning-go-lang/functions"
	"github.com/whitneylampkin/learning-go-lang/variables"
)

func main() {
	fmt.Println("Hello World!")

	// Referencing a local package
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

	// Referencing 3rd-party packages
	fmt.Println(quote.Hello())

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
	sum := functions.CustomSumFunction(os.Args[1], os.Args[2])
	println("Sum with a custom function:", sum)

	// Calling returnMultipleValues()
	sum, mul := functions.ReturnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	// Discarding variables
	sum2, _ := functions.ReturnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum2 (Discarding mul):", sum2)

	functions.PointerParametervalues()
}
