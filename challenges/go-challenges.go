package challenges

import "fmt"

func FindPrimes(number int) bool {
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

func Fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}

func PanicChallenges() {
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
	}
}

func CalculateFibonacci() {
	// Get users input
	userInput := 6

	if userInput < 2 {
		fmt.Print("\r\n")
		fmt.Println(make([]int, 0))
	}

	fibonacciSequence := make([]int, userInput)

	fibonacciSequence[0], fibonacciSequence[1] = 1, 1

	for i := 2; i < userInput; i++ {
		fibonacciSequence[i] = fibonacciSequence[i-1] + fibonacciSequence[i-2]
	}

	fmt.Print("\r\n")
	fmt.Println(fibonacciSequence)
}

/* func TranslateRomanNumerals() {
	romanNumbers := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// Get user input
	userInput := "MCLX"

	var sum int

	// Sum user input
	for i := 0; i < len(userInput); i++ {
		value, exist := romanNumbers[userInput[i]]

		if exist {
			sum += value
		} else {
			panic("Invalid roman numeral value entered.")
		}
	}

	fmt.Println("Roman Number Sum: ", sum)
} */
