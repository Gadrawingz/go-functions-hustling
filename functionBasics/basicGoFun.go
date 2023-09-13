package main

import (
	"fmt"
)

/*
LOGIC:
Functions used  divide our code into smaller chunks to make our code looks clean and easier to understand.

Basically, a function is a block of code that performs a specific task.
For example, suppose we want to write code to create a circle and rectangle and color them.

In this case, we can organize our code by creating three different functions:
	-function to create a circle
    -function to create a rectangle
    -function to color

This way, Our code will look more organized. Also, we can reuse the same
function to color the circle and rectangle. Hence, providing the benefits of
code reusability.
*/

func addNumbers(number1 int, number2 int) {
	var sum = number1 + number2
	fmt.Println(sum)
}

/*
Function Parameter
------------------
However, in real projects, we want our functions to work dynamically. That is, instead of adding 12 and 8, the addNumbers() function should be able to add any two numbers.

In this case, we can create functions that accept external values and perform operations on them. These additional parameters are known as function parameters
*/

// str1 and str2 are parameters
// the values will passed are known as function arguments:
func mergeWords(str1 string, str2 string) {
	result := str1 + "" + str2
	fmt.Println(result)
}

/*
If consecutive parameters are of the same type, we can avoid writing
the type each time and it is enough to be written once at the end.ie
price int, no int can be written as price, no int.
Example: func calculateBill(price, productNo int) int {
*/
func calculateBill(price int, productNo int) int {
	var totalPrice = price * productNo
	return totalPrice
}

// Multiple return values (1):
func rectangleProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length * width) * 2
	return area, perimeter
}

// Multiple return values (2):
func sumAndSub(num1 int, num2 int) (int, int) {
	mySum := num1 + num2
	var mySub = num1 - num2
	return mySum, mySub
}

// Lets return multiple values of different data types
func getDefaultValues() (int, string) {
	return 0, "Not Known"
}

// Names return values
func rectanglePropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	// No explicit return value...
	return
}

// Using Blank identifier (_)
// Used if we only need the area and want to discard the parameter
func getRectangleProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func sayHello() {
	fmt.Println("Good Morning!")
}

func main() {
	fmt.Println("********************************************************")
	sayHello() // Function function

	// Calling named function
	fmt.Println(rectanglePropsNamed(10, 4))

	// Calling a function with Blank identifier
	myArea, _ := getRectangleProps(20.4, 3.2)
	fmt.Printf("The area is: %f \n", myArea)

	// Calling sumAndSub function
	var mySum, mySub = sumAndSub(50, 5)
	fmt.Println("The sum is: ", mySum, "The sub is: ", mySub)

	// Calling n using function with the default value!
	myId, firstname := getDefaultValues()
	fmt.Println("ID: ", myId, ", First Name: ", firstname)
	fmt.Println("************************************************")

	// 1st Usage:
	addNumbers(34, 66)
	addNumbers(4222, 78)

	// 2nd Usage: (A)
	fmt.Println("The calculated value:  ", calculateBill(400, 10))

	// 2nd Usage: (B)
	thePrice, prodNo := 50, 10
	theTotalPrice := calculateBill(thePrice, prodNo)
	fmt.Println("The new price is :  ", theTotalPrice)

	fmt.Println("********************************************************")

	// Another function call:
	mergeWords("Spider", "Man")
	mergeWords("Iron", "Man")

}
