package main

import (
	"fmt"
)

// const a = 42
// const b = 42.78
// const c = "James Bond"
// alternatively

const (
	a = 42
	b = 42.78
	c = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}

// Constants are simple, unchanging values. They only exist at compile time
// There are TYPED and UNTYPED constants (const hello = "Hello") (const typedHello string = "Hello")
// Untyped constants do not yet have a fixed type, they grant us some flexibility with the compiler
// Type isn't chosen until compilation (is type int, uint, float64)? Avoids needing to do specific conversion
