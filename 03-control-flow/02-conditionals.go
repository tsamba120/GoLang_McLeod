package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("hello")
	}

	if 2 != 2 {
		fmt.Println("this will not print")
	}

	if !(3 == 4) {
		fmt.Println("print this")
	}

	// use of initialization statement with colon
	if x := 42; x != 2 {
		fmt.Println("statement")
		fmt.Println(x)
	}

	// if + else; if + elseif + else
	y := 42
	if y == 40 {
		fmt.Println("our value was 40")
	} else if y == 100 {
		fmt.Println("our value was 41")
	} else {
		fmt.Printf("value was not 40 or 41. actual value was %d\n", y)
	}

	fmt.Println()
	// CONDITIONAL LOGIC OPERATORS
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
