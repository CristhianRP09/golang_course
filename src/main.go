package main

import "fmt"

func normalFunction(msg string) {
	fmt.Printf(">>>>>>>>>>> Hello world %s\n", msg)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func duplicate(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, 2 * a
}

func main() {
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	name := "Cristhian"
	age := 27
	fmt.Printf("My name is %s and I am %d years old\n", name, age)
	fmt.Printf("My name is %v and I am %v years old\n", name, age)

	// Sprintf
	message := fmt.Sprintf("My name is %s and I am %d years old", name, age)
	fmt.Println(message)

	// Data type
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("name: %T\n", name)

	// Functions
	normalFunction("1")
	normalFunction("2")
	tripleArgument(1, 2, "Hello!")

	value := duplicate(3)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value1:", value1)
	fmt.Println("Value2:", value2)

	// Just using the first returned value
	v1, _ := doubleReturn(5) // _ means ignoring the 2nd value in this case
	fmt.Println("v1:", v1)
}
