package main

import "fmt"

func main() {

	//   Remember:
	// - In Go, a string is a read-only slice of bytes
	// - Strings in Go are UTF-8 encoded by default
	// - all characters are represented in int32 (size of 4 bytes) data type because
	//   UTF-8 chars can be rappresented from 1 byte to 4 bytes
	// - UTF-8 code point or character value can be represented by
	//   series of one or more bytes (max 4 bytes)

	//==================
	// An ASCII string
	//===================
	s := "Hello world"

	// result will be
	// 72   101  108  108  111  32  119  111  114  108  100
	//  these are the decimal value of ASCII/UTF-8 characters in Hello World string
	//  we are seeing the byte (uint8) values of string s which is internally a slice
	//  Hence s[i] prints the decimal value of the byte held by the character.

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], " ")
	}

	fmt.Println(" ")
	//==================
	// An string with NON ASCII chars
	//===================
	na := "Hellõ world"

	// result:
	//72  101  108  108  195  181  32  119  111  114  108 100
	// Note that the number of bytes is 12, so if we do
	// len(na) we will get 12. We have a byte more than the
	// previous because the char "õ" occupies two bytes for its
	// code point

	for i := 0; i < len(na); i++ {
		fmt.Println(na[i], " ")
	}

	fmt.Println(" ")

	//..but if we try to see chars with a for and indexing the string
	// we will see different chars becuase indexing the stirng we get only
	// one byte

	for i := 0; i < len(na); i++ {
		fmt.Printf("%c", na[i])
	} // ... HellÃµ world

	fmt.Println(" ")

	// to avoid this, we have "rune" data type (synonim of code point).
	// - rune is an alias of int32

	//..so we need to convert the string into a slice of runes:
	r := []rune(na)

	// so now we can see the right result
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	fmt.Println(" ")

	//===============
	// USING RANGE
	//=================

	//range will return rune and byte index of the character.
	//note: we lost index 5 because the 5th byte is second code unit of õ character
	for index, char := range na {
		fmt.Printf("character at index %d is %c\n", index, char)
	}
	//character at index 0 is H
	//character at index 1 is e
	//character at index 2 is l
	//character at index 3 is l
	//character at index 4 is õ
	//character at index 6 is
	//character at index 7 is w
	//character at index 8 is o
	//character at index 9 is r
	//character at index 10 is l
	//character at index 11 is d

}
