package main

import (
	"fmt"
)

var a int

type hotdog int // created our own type "hotdog" with underlying type as int
var b hotdog    // declared variable b of type hotdog

func main() {
	a = 42
	b = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b) // returns "main.hotdog" - type hotdog from package main

	// a = b -- we CANNOT do this because a and b are assigned two different types
	// even though b has the same underlying type as a
}
