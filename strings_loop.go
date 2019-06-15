package main

import "fmt"

func main() {

	s := "Hello world!"

	fmt.Println(len(s))

	// result will be
	// 72   101  108  108  111  32  119  111  114  108  100
	// these are the decimal value of ASCII/UTF-8 characters in Hello World string
	//  we are seeing the byte (uint8) values of string s which is internally a slice
	// Hence s[i] prints the decimal value of the byte held by the character.

	// remember:
	// - In Go, a string is a read-only slice of bytes
	// - Strings in Go are UTF-8 encoded by default

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], " ")
	}

	fmt.Println(" ")
	// If you want to see chars you can use fmt.Printf with  %c
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}

	fmt.Println(" ")

	// If you want to see byte value you can use fmt.Printf with  %v
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	fmt.Println(" ")

	// If you want to see hex value you can use fmt.Printf with  %v
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println(" ")
}
