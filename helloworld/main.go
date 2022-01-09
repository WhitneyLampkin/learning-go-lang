package main

import (
	"fmt"
	"learning-go-lang/calculator"
	"learning-go-lang/controlflows"
	"learning-go-lang/datatypes"
	"learning-go-lang/forloops"
	"learning-go-lang/functions"
	"learning-go-lang/variables"
	"os"
)

func main() {
	fmt.Println("Hello World!")

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

	// Forloops
	forloops.BasicForLoop()
}
