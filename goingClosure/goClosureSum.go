package main

import "fmt"

/*
Go functions may be closures.
=============================
=> A closure is a fx value that references variables from outside its body.
=> The fx may access and assign to the referenced variables; in this sense
the fx is "bound" to the variables.

The adder fx returns a closure. Each closure is bound to its own sum variable.
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	// Going!
	pos, neg := adder(), adder()

	// Looping!
	for i := 0; i < 5; i++ {
		fmt.Println("POS: ", pos(i), "NEG:", neg(-2*i))
	}
}
