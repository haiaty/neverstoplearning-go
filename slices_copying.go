package main

import "fmt"

func main() {

	//create a new empty slice
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	//creates a new slice with the dimension
	// of the previous one
	c := make([]string, len(s))

	copy(c, s)
	fmt.Println("cpy:", c) //cpy: [a b c]

}
