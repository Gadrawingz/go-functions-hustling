package main

import "fmt"

/*
=> Go provides a special feature known as an "anonymous fx".
=> An anonymous function can form a closure.

=> A closure is a special type of anonymous function that references variables declared outside of the function itself.
=> It's similar 2 accessing global variables which are available b4
the declaration of the fx.

=> Go closure is a nested function that allows us to access variables of the
outer function even after the outer function is closed.

Before doing closure, we must understand:
=========================================
-> 1. NESTED FX: It is a created function inside another function.
-> 2. Returning a function in Go
*/

// 1A. NESTED FX: It is a created function inside another function.
// ===============================================================

// Outer fx
func greetPerson(person string) {

	// Inner fx
	var displayGreetedPerson = func() {
		fmt.Println("Hello", person)
	}

	// Call inner fx
	displayGreetedPerson()
}

// 1B. Returning a function in Go
// ===============================================================
// func greetSomeone() func() {...}, func() b4 the {} indicates that the fx
// returns another function.
// If U look into the return statement of this fx, the fx is returning a fx.
func greetSomeone() func() {

	return func() {
		fmt.Println("Hello Some1!")
	}
}

func main() {

	// Calling outer fx (1)
	greetPerson("Bruno")
	greetPerson("Joyner Lucas")

	// Calling (2)
	greeting1 := greetSomeone()
	greeting1()
}
