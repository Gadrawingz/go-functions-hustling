package main

import "fmt"

// A simple go program that generate a sequence of integers.
// Using intSequence fx. It returns a closure which increments the m variable.
// Simple: The intSequence is a fx w/c returns a fx that returns an integer.

func intSequence() func() int {

	m := 0

	return func() int {
		m++
		return m
	}
}

func main() {

	// We can call the closure several times

	// The first call
	theNextInteger1 := intSequence()
	fmt.Println(theNextInteger1()) // 1
	fmt.Println(theNextInteger1()) // 2
	fmt.Println(theNextInteger1()) // 3
	fmt.Println(theNextInteger1()) // 4
	fmt.Println("--------------------")

	// The next call of the intSequence fx returns a new closure.
	// This new closure has its own distinct state
	theNextInteger2 := intSequence()
	fmt.Println(theNextInteger2()) // 1
	fmt.Println(theNextInteger2()) // 2
}
