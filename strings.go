package main

import "fmt"

func main() {

	// - In Go, a string is a read-only slice of bytes
	// - Strings in Go are UTF-8 encoded by default
	mystring := "Hi"

	fmt.Println(mystring) // Hi

	fmt.Println(len(mystring)) // 2

}
