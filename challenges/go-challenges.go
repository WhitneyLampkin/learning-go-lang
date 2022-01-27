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
