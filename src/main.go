package main

import (
	"fmt"
	mypkg "golang_course/src/mypackage"
)

func main() {
	// Interfaces
	mySquare := mypkg.Square{Base: 10}
	myRectangle := mypkg.Rectangle{Base: 5, Height: 8}

	// Without interfaces, we need to run the area methods for each instance:
	fmt.Println("Square area:", mySquare.Area())
	fmt.Println("Rectangle area:", myRectangle.Area())

	// With interfaces:
	mypkg.ComputeArea(mySquare)
	mypkg.ComputeArea(myRectangle)

	// List of Interfaces
	mixedList := []interface{}{"Hey!", true, 123, -3.1324}
	fmt.Println(mixedList)
}
