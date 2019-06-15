package main

import "fmt"

func main() {

	//create a new empty slice
	s := make([]string, 3)
	fmt.Println("emp:", s) // emp: [  ]

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    //set: [a b c]
	fmt.Println("get:", s[2]) // get: c

	fmt.Println("len:", len(s)) //3

	//=========================
	// COPYING SLICES
	//=========================

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) //cpy: [a b c d e f]

}
