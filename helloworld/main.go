package main

import (
	"fmt"
	"os"

	"github.com/whitneylampkin/learning-go-lang/arrays"
	"github.com/whitneylampkin/learning-go-lang/calculator"
	"github.com/whitneylampkin/learning-go-lang/challenges"
	"github.com/whitneylampkin/learning-go-lang/controlflows"
	"github.com/whitneylampkin/learning-go-lang/datatypes"
	"github.com/whitneylampkin/learning-go-lang/errorhandling"
	"github.com/whitneylampkin/learning-go-lang/fizzbuzz"
	"github.com/whitneylampkin/learning-go-lang/forloops"
	"github.com/whitneylampkin/learning-go-lang/functions"
	"github.com/whitneylampkin/learning-go-lang/logging"
	"github.com/whitneylampkin/learning-go-lang/maps"
	"github.com/whitneylampkin/learning-go-lang/methods"
	"github.com/whitneylampkin/learning-go-lang/slices"
	"github.com/whitneylampkin/learning-go-lang/structs"
	"github.com/whitneylampkin/learning-go-lang/variables"
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
	controlflows.DeferExample()
	// The next line is commented out because it's stops execution.
	// controlflows.PanicHighLowExample()
	controlflows.RecoverExample()

	// Forloops
	forloops.BasicForLoop()
	forloops.FakeWhileLoop()
	forloops.InfiniteLoop()

	// Challenge: FizzBuzz
	fizzbuzz.Fizzbuzz()

	// Challenge: Find the primes (under 20)
	fmt.Println("Prime numbers less than 20:")

	for number := 2; number < 20; number++ {
		if challenges.FindPrimes(number) == true {
			fmt.Printf("%v ", number)
		}
	}

	// Arrays
	arrays.DeclareArrays()
	arrays.InitializingArrays()
	arrays.EllipsisInArrays()
	arrays.ArrayLastPositionOnly()
	arrays.MultidimensionalArrays()

	slices.DeclareAndInitializeASlice()
	slices.SliceItems()
	slices.ExtendedSlice()
	slices.AppendSliceItems()
	slices.RemoveSliceItems()
	slices.CopySliceItems()
	slices.CopySliceItemsFixed()

	maps.DeclareAndInitializeMaps()
	maps.AddMapItems()
	// Uncomment next line to see the error message
	// maps.CreateFailedAddMapItemsCase()
	maps.AccessMapItems()
	maps.RemoveMapItems()
	maps.LoopsThroughMaps()

	structs.DeclareAndInitializeStructs()
	structs.EmbeddingStructs()
	structs.EncodeDecodeStructsJSON()

	challenges.CalculateFibonacci()
	// challenges.TranslateRomanNumerals()

	// Error handling
	errorhandling.HandleError()

	// Logging
	logging.LogStandardMsg()
	// logging.LogFatalMsg()
	// logging.LogPanicMsg()
	//logging.LogSetPrefix()
	logging.LogToFile()

	// Methods
	methods.DeclareMethod()

	// Interfaces

	// Concurrency
}
