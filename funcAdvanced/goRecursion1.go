package main

import (
	"fmt"
)

/*
1. Recursive fx call itself.
============================
Here, Recurse fx includes call within its body.
// The technique is called recursion
// EX: func recurse() { .... recurse() }

EXAMPLE1: FX1: Making countDown()
Ref:https://www.programiz.com/golang/recursion
*/

func countDownFx(number int) {

	if number > 0 {
		fmt.Print(number, " ")

		// Recursive call is added inside if statement by decreasing number
		countDownFx(number - 1)
	} else {
		// Ends the recursive fx
		fmt.Println("CountDown stops!\n")
	}
}

func main() {

	// Function call FX1:
	countDownFx(14)
	countDownFx(7)
	countDownFx(25)

}
