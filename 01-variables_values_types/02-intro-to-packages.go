package main

import (
	"fmt"
)

func main() {
	// fmt.Print can also return (byte size, error code) tuple that we can unpack below
	n, _ := fmt.Println("Hello, playground", 42, true) // "_" throws away value
	fmt.Println(n)

	_, e := fmt.Println("Hello, playground", 42, true)
	fmt.Println(e)
}

// VARIADIC FUNCs
// Functions in general accept only a fixed number of arguments.
// fmt.Println is an example of a variadic function
// such a function accepts a variable number of arguments

// INTERFACE TYPES
// type "interface{}" - we see this as a parameter in variadic functions
// means "give me as many arguments of any types as you'd like"

// THROWING AWAY RETURNS
// use the "_" underscore character to throwaway returns
// you can't have unused variables in your code. this is code pollution and code won't compile
// if instead we did 'n, e := fmt.Println("Hellow....") -> not compile

// IDENTIFYING GO PACKAGES
// package.Identifier
// ex: fmt.Println - "from package fmt, we use the Println func"
