package main

import "fmt"

func main() {

	// Just declaring my global variable
	theNumber := 30

	// Assigning an anonymous fx to a variable
	myCounter := func() int {

		theNumber += 1
		return theNumber
	}

	fmt.Println(myCounter()) // 31
	fmt.Println(myCounter()) // 32

	/*
		Note1: The variable theNumber wasn't passed as a parameter 2the anonymous fx
		but the function has access to it.

		Note2: A slight problem as any other fx which will be defined in the
		main has access to the global variable theNumber and it can change it
		without calling counter fx.

		Sol: Thus closure also provides another aspect which is data isolation.
		Ref: newCounter is created below!
	*/

	// newCounter fx is assigned to a variable
	theCounter := newCounter()

	// Invoke the counter
	fmt.Println(theCounter()) // 21
	fmt.Println(theCounter()) // 22
	fmt.Println(theCounter()) // 23

}

// newCounter fx to isolate global variable
func newCounter() func() int {
	theNumber := 20
	return func() int {
		theNumber += 1
		return theNumber
	}
}

/*
The closure references the variable "theNumber" even after the newCounter() fx
has finished running but no other code outside of the newCounter() fx has access  to this variable.

This is how data persistency is maintained between function calls while
also isolating the data from other code.
*/
