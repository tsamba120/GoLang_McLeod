package main

import "fmt"

// ZERO VALUE
// A zero value is the "default" value a variable takes when it is initialized and given a type,
// BUT the variable is not given a value
// i.e. strings default value is "", int default value is 0, boolean default value is false, float default is 0.0
// note it is "nil" for things like pointers, maps, etc.

var x string
var y int
var z bool
var a float32

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", a)
}
