package main

import "fmt"

func main() {

	// Any valid UTF-8 character within a single quote (')
	// is a rune and it’s type is int32
	r := 'õ'

	//print char
	fmt.Printf("%c", r) //õ

	fmt.Println(" ")

	// print the byte value
	fmt.Printf("%v", r) // 245

	fmt.Println(" ")

	// print the hex value
	fmt.Printf("%x", r) //f5

	fmt.Println(" ")

	// print datatype
	fmt.Printf("%T", r) //int32

	fmt.Println(" ")

}
