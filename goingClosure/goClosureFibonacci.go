package main

import "fmt"

/*
MAKING A FIBONACCI SERIES
=========================
Fibonacci series is a sequence of values such that each number is the sum of 
the two preceding ones, starting from 0 and 1. The beginning of the sequence 
is thus: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 ...
*/

func getFibonacci() func() int {

	var a = 0
	var b = 1
	return func() int {
		a, b = b, a+b
		return b - a
	}

	Fibonacci series is a sequence of values such that each number is the sum of the two preceding ones, starting from 0 and 1. The beginning of the sequence is thus: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 ...

}

func main() {

	fibonacci1 := getFibonacci()
	for i := 0; i < 7; i++ {
		fmt.Print(fibonacci1(), " ")
	}
	fmt.Println() // 0 1 1 2 3 5 8

	fibonacci2 := getFibonacci()
	for i := 0; i < 14; i++ {
		fmt.Print(fibonacci2(), " ")
	}
	fmt.Println() // 0 1 1 2 3 5 8 13 21 34 55 89 144 233

}
