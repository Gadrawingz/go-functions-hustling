package main

import "fmt"

// Go program to calculate the sum of positive numbers!
func sumUp(num int) int {
	// Let's Condition to break recursion
	if num == 0 {
		return 0
	} else {
		result := num + sumUp(num-1)
		return result
	}
}

func main() {

	// Function calls:
	fmt.Println("SUM: ", sumUp(5))  // 1+2+3+4+5 => 15
	fmt.Println("SUM: ", sumUp(10)) // 1+2+3+4+5+6+7+8+9+10 => 55

	// Other way to call
	var anyNumber1 = 4
	var anyNumber2 = 6
	var result1 = sumUp(anyNumber1)
	var result2 = sumUp(anyNumber2)
	fmt.Println(result1) // 1+2+3+4 => 10
	fmt.Println(result2) // 1+2+3+4+5+6 => 21

}
