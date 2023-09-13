package main

import (
	"fmt"
)

// 2. Global variables
// When we declare variables before the main() function, these variables
// will have global scope. We can access them from any part of the program.
var outSum int

// 1. Local variables:
// Are those type of variables declared inside a function, which means these
// will have a local scope (within fx). We cannot access them outside the fx
func printNumbers() {
	// local variables
	var number1 int
	var number2 int
	var number3 int
	var inSum int

	number1, number2, number3 = 40, 7, 20
	inSum = number1 + number2 + number3
	fmt.Println(inSum)
}

func printNumbersUseGlobal() {
	// local variables
	var number1 int = 35
	var number2 int = 62
	// Use our global variable
	outSum = number1 + number2
}

func main() {

	// Calling F1 function
	printNumbers() // 67

	// Lets use global variable;
	fmt.Println("OUT SUM:", outSum) // OUT SUM: 0

	// But
	printNumbersUseGlobal()
	fmt.Println("OUT SUM:", outSum) // OUT SUM: 97

}
