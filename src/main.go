package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Is Palindrome!")
	} else {
		fmt.Println("Is Not Palindrome!")
	}
}

func main() {
	// Arrays
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(cap(array))

	// Slices
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Slice methods
	fmt.Println(slice[0])   // Shows the 1st element of slice
	fmt.Println(slice[:3])  // Shows the from 0 to 2nd elements of slice
	fmt.Println(slice[2:4]) // Shows the from 2 to 3rd element (without 4th) of slice
	fmt.Println(slice[4:])  // Shows elements from 4th onwards

	// Append Method
	slice = append(slice, 7) // Add element 7 to slice
	fmt.Println(slice)
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Add element the slice {9, 9, 10} to previous slice
	fmt.Println(slice)

	// Range method
	otherSlice := []string{"Hey", "How", "are", "you"}

	for i, value := range otherSlice {
		fmt.Println(i, value)
	}

	// Exercise
	isPalindrome("ama")         // true
	isPalindrome("amor a roma") // true
	isPalindrome("casas")       // false
	isPalindrome("amA")         // true
	isPalindrome("aMoR A ROmA") // true
}
