package main

import "fmt"

func main() {

	/*
		1. ANONYMOUS FUNCTION:
		-> It is function without the function name.
		-> It works just like a regular function in Go. Syntax:
		func () {
			fmt.Println("Function without name")
		}

	*/

	// 1. Working of Go Anonymous Function (1)

	// Usually, We assign the anonymous function to a variable and
	// then use the variable name to call the function.

	// Anonymous function
	var greetPeople = func() {
		fmt.Println("Good Evening!")
	}

	// Anonymous function call
	greetPeople() // Good Evening!
	// This is not the way we call it!
	fmt.Println(greetPeople) //0x481880

	// 2. Anonymous Function with Parameters
	var makeAddition = func(num1 int, num2 int) {
		add := num1 + num2
		fmt.Println("SUM: ", add)
	}
	makeAddition(54, 46) // SUM: 100
	makeAddition(14, 27) // SUM: 41

}
