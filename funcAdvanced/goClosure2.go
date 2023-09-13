package main

import "fmt"

// Its time for closure
// Closure is a nested fx that helps us access outer fx's
// variable even after fx is closed

func congratulate() func() string {

	// var defined outside the inner fx
	username := "Stephen"

	// return a nested anonymous fx
	return func() string {
		username = "Congrats, " + username + "!!"
		return username
	}
}

func main() {

	// Call the outer fx
	message := congratulate()

	// Call the inner fx
	fmt.Println(message())

}
