package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Conditional statements
	value1 := 1
	value2 := 3

	if value1 == 1 {
		fmt.Println("Equals to 1")
	} else {
		fmt.Println("Not equal to 1")
	}

	// AND operator
	if value1 == 1 && value2 == 2 {
		fmt.Println("TRUE")
	}

	// OR operator
	if value1 == 0 || value2 == 3 {
		fmt.Println("TRUE")
	}

	//  Parsing text into numbers
	value, error := strconv.Atoi("53")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Value:", value)

	// Conditional for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// For while loop
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever loop
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
