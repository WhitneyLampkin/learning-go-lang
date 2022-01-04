package calculator

/*
	Go doesn't have public and private keywords

	private - start with lowercase letter
	public - start with uppercase letter
*/

// private variable
var logMessage = "[LOG]"

// Version of the calculator
// public variable
var Version = "1.0"

// private function
func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
// public function
func Sum(number1, number2 int) int {
	return number1 + number2
}
