package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the value the string "Shaken, not stirred"
// REMEMBER go is a static programming language, not dynamic
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
var z string = "Shaken, not stirred"

// raw string literal - backticks allow inclusion of internal double quotes and new lines as seen below
var a string = `James said, "Shaken, 

not stirred"`

func loop_printer(limit int) {
	total := 0
	for i := 0; i <= limit; i++ {
		fmt.Println(total)
		total++
	}
}

func main() {
	fmt.Println(y)
	fmt.Printf("The variable 'y' is of type: %T\n", y) // %T prints data type of an object
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// z = 42 -> this is NOT possible as we declared z as an integer variable
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

// STRING LITERALS
// Raw string literals are defined inside back ticks, i.e. `...`
// Printing a raw string literal will return anything between these back ticks, include double quotes
