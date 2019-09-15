package main

import "fmt"

func main() {

	// - In Go, a string is a read-only slice of bytes
	// - Strings in Go are UTF-8 encoded by default
	mystring := "Hi"

	fmt.Println(mystring) // Hi

	// Remenber: a string is a read-only slice of bytes
	//  When we use len function on a string,
	//  it calculates the length of that slice
	fmt.Println(len(mystring)) // 2

	//Strings are immutable
	mystring[0] = "j" //compile error: cannot assign to mystring[0]

}
