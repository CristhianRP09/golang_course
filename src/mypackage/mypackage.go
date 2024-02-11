package mypackage

import "fmt"

// Structs defined with first character in lowercase are private and
// those defined with capitalized names are public. Same applies
// for attributes within the struct.

// CarPublic - Example of public struct
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate - Example of private struct
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage - Example of public function
func PrintMessage(msg string) {
	fmt.Println(msg)
}

// printMessagePrivate - Example of private function
func printMessagePrivate(msg string) {
	fmt.Println(msg)
}
