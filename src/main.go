package main

import (
	"fmt"
	mypkg "golang_course/src/mypackage" // mypkg is an alias for this package
)

// Defining a struct
type car struct {
	brand string
	year  int
}

func main() {
	// Classes in Golang: Structs

	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Another way to instantiate an struct
	var anotherCar car
	anotherCar.brand = "Ferrari"
	fmt.Println(anotherCar) // non-specified attributes (year) are set up to zero-value!

	// Importing from other packages
	var vehicle mypkg.CarPublic // Instantiates a CarPublic struct
	vehicle.Brand = "Ferrari"
	vehicle.Year = 2020
	fmt.Println(vehicle)

	mypkg.PrintMessage("Hey you!") // Executes the function defined in mypkg package!
}
