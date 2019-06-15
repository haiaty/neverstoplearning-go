package main

import "fmt"

func main() {

	//create a new empty slice
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	//note that we get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) //apd: [a b c d e f]

	//=========================
	// COPYING SLICES
	//=========================

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) //cpy: [a b c d e f]

}
