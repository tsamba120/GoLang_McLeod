package main

import (
	"fmt"
)

func main() {
	// format
	// for init; condition; post {}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("While loop:")

	// while loop equivalent
	t := 0
	for t < 5 {
		fmt.Println(t)
		t++
	}

	fmt.Println("Infinite loop with break:")
	// infinite loop with break
	t = 0
	for {
		if t > 9 {
			break // continue keyword also exists
		}
		fmt.Println(t)
		t++
	}
}

// NOTE: you can use the break statement like all other languages,
// however if you are inside nested logic and you want to exit out of an external loop,
// use the 'break Loop' keyword to specify breaking out of a loop and not something like a switch statement
// we'll learn about the "for range" mechanism later on
