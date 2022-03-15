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

	a = int(b) // converts var b to type "int" and stores in var a which is declared as type "int"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

// NOTE
// In Go we refer to this as "converting" and not "casting"
// Similar to "casting" in other language we can only "convert" variables that CAN be converted
// Ex: we can't convert a variable valued as "hello" to an int or boolean
