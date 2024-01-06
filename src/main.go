package main

import "fmt"

func main() {
	// Declaring constants
	const pi float64 = 3.141592
	const pi2 = 3.1416

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaring integer variables
	base := 12          // := declares and initializes a variable
	var height int = 14 // declare its type and assing a value
	var area int        // declares its type

	fmt.Println(base, height, area)

	// Zero values
	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false

	fmt.Println(a, b, c, d)

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area", squareArea)

	// Aritmetic operations
	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("Sum: ", result)

	// Diff
	result = y - x
	fmt.Println("Difference: ", result)

	// Product
	result = x * y
	fmt.Println("Product: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Module
	result = y % x
	fmt.Println("Module: ", result)

	// Increment
	x++
	fmt.Println("Increment: ", x)

	// Decrement
	x--
	fmt.Println("Decrement: ", x)
}
