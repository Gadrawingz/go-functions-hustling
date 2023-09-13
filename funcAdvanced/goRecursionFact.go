package main

import "fmt"

// Finding factorial number using recursion
func factorial(number int) int {

	// condition 2break recursion
	if number == 0 {
		return 1
	} else {
		var result = number * factorial(number-1)
		return result
	}
}

// 6 * 5

func main() {

	// Function call & Printing!
	// 120 = 5 * 4 * 3 * 2 * 1
	fmt.Println("FACT OF 5:", factorial(5))

	// 6 = 3 * 2 * 1
	fmt.Println("FACT OF 3:", factorial(3))

	// 720 = 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println("FACT OF 6:", factorial(6))
}
