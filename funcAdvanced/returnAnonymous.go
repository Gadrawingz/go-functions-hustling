package main

import (
	"fmt"
)

// Ref: https://www.programiz.com/golang/anonymous-function
// A Regular function to calculate square of numbers
func findSquare(num int) int {
	square := num * num
	return square
}

func main() {

	// 1. Return Value From Anonymous Function
	// A. Mhttps://www.programiz.com/golang/anonymous-functionake a function that return the division

	var divide = func(largeNum, smallNum int) int {
		divide := largeNum / smallNum
		return divide
	}

	// A. Function call
	result1 := divide(40, 2)
	fmt.Println("DIVISION:", result1) // DIVISION: 20
	fmt.Println(divide(73, 6))        // 12

	// B. Program to return area of a rectangle
	area := func(length, breadth int) int {
		return length * breadth
	}
	fmt.Println(area(30, 20)) // 600
	fmt.Println(area(54, 12)) // 648

	// 2. Passing anonymous function as arguments to other functions
	// We have a func to find square, we want to pass anonymous function for
	// calculation sum inside a function that find square root
	var getSum = func(n1, n2, n3 int) int {
		return n1 + n2 + n3
	}

	// Function call!
	theResult := findSquare(getSum(40, 3, 7))
	fmt.Println(theResult) // 2500

	// Which means:
	theSum := getSum(40, 3, 7)     // 50
	theResult = findSquare(theSum) // 50 x 50
	fmt.Println(theResult)         // 2500 (Result)
	fmt.Println("**************************************")

	// Return an anonymous Function in Go
	// R1 Func calling! (Ref: its declared down below)
	salary1 := printSalaryWithBonus()
	fmt.Println(salary1()) // 65000
	fmt.Println(salary1()) // 80000
	fmt.Println(salary1()) // 95000

	salary2 := printSalaryWithBonus()
	fmt.Println(salary2()) // 65000

	// R2 Function call
	justNumber := displayGrowingNumber()
	fmt.Println(justNumber()) // 51.5

}

// Return an anonymous Function in Go
// 4that, We create an anonymous fx inside a regular fx and return it.
// C. Go Program to return an anonymous function
// R1 Func declaration
func printSalaryWithBonus() func() int {
	salary := 50000
	// Anonymous fx inside
	return func() int {
		salary += 15000 // With extra 15000
		return salary
	}
}

/*
Here, we're returning the anonymous fx. So, when the printSalaryWithBonus()
function is called, the anonymous function is also called and the incremented number is returned.
*/

// R2: Another function
func displayGrowingNumber() func() float64 {
	number := 50.5
	return func() float64 {
		number++
		return number
	}
}

// Do
