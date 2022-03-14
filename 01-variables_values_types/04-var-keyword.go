package main

import "fmt"

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with IDENTIFIER "z" is of TYPE int
var z int = 105

// NOTE! We can use "var" when we need to declare a variable outside of func body
// the scope of z is the entire program, or "package level"
// not possible to use short declaration operator outside of func body

var a int // defaults to zero value

func main() {
	// short declaration operator
	// DECLARE	a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	var y int = 42 // var y = 42 alternatively
	fmt.Printf("This is the number forty-two, %d\n", y)

	foo()
}

// we can declare a function to print global variable z
func foo() {
	fmt.Println(z)
	fmt.Println("Default value for int type:", a)
}

// BEST PRACTICE:
// limit the scopre of variables
// as much as possible try to use short declaration operator (only possible inside func body)

// DEFAULT VALUES:
// booleans -> false
// integers -> 0
// floats -> 0.0
// strings -> ""
// pointers, functions, interfaces, slices, channels, and maps -> nil
