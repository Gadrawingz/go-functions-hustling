package main

import "fmt"

/*
1. Return Value from Go Function
---------------------------------
We can return value from a function and use it anywhere in our program
*/

// EX1 function
func addTwoNumbers(n1 float64, n2 float64, n3 float64) float64 {
	var sumNumbers = n1 + n2 + n3
	return sumNumbers
}

// EX2: The return statement should be the last statement of function
// When a return statement is encountered, the function terminates.
// Below We get a "missing return at the end of function" error as
// we have added a line after the return statement.

/*
func add(x int, y int) int {
	sum := x + y
	return sum
	// code after return statement
	fmt.Println("After return statement")
}
*/

// ============================================
// 2. Return Multiple Values from Go Function
// In Go, we can also return multiple values from a function.

// (int, int) indicates we're returning 2integer values:sum&diff 4rom the fx
func calculateSD(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2

	// return 2 values;
	return sum, difference
}

func main() {

	// EX1 function call (1)
	fmt.Println(addTwoNumbers(3.5, 12, 4.5)) // 20

	// Ex1 function call (2)
	myResult1 := addTwoNumbers(40, 21.4, 8.6)
	fmt.Println(myResult1) // 70

	// Ex3 function call (1)
	var theSum, theDiff = calculateSD(30, 24)
	fmt.Println("SUM:", theSum, "DIFF: ", theDiff) // SUM: 54 DIFF:  6

	var _, theDiff2 = calculateSD(43, 15)
	fmt.Println("THE DIFF: ", theDiff2) // THE DIFF: 28 (sum was ignored)
	// Ex3 function call (2)

}
