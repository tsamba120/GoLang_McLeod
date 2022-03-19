package main

import (
	"fmt"
)

// Within a constant declaration, the predeclared identifier "iota" represents successive untyped integer constants.
// It is reset to 0 whenver the reserved word "const" appears in the source and increments after each ConstSpec
// Auto-increments whenever called

const (
	a = iota
	b = iota
	c = iota
)

const ( // iota will reset here. also: another way to declare
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println("\n'const' keyword is called so iota resets to 0 and increments again")
	fmt.Println()

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
