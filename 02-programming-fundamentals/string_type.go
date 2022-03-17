package main

import (
	"fmt"
)

// as with string data types in other languages - strings are immutable
// you can reassign a variable a new string but you can't make changes to the value in-place
func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	t := `"hello"`
	fmt.Println(t)

	// Strings are "slices of bytes"
	bs := []byte(s) // this casts a string into a slice of bytes (note ascii)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// This for loop iterates through each byte (character) in the string and prints its UTF-8 code point
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	// Prints index and byte value + hex value
	for i, v := range s {
		fmt.Printf("Index %d: byte value = %d and hex value = %#x\n", i, v, v)
	}

}
