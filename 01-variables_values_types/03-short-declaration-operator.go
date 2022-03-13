package main

import "fmt"

func main() {
	a := 100 // declares AND assigns variable
	fmt.Println("Newly declared and assigned variable 'a', uses ':='", a)
	a = 5 // assigns a NEW value to already declared variable. only need to use equal sign
	fmt.Println("Reassigned already-declared variable 'a', uses '=':", a)
	// a = "hello" - this line will cause line to NOT compile because variable already declared as an int
	// Remember GO is strongly, statically typed like Java

	x := 42
	y := 24
	fmt.Println("\n", x == y)

	// Sprintf returns a formatted string
	jackieRobinsonReponse := fmt.Sprintf("Jackie Robinson's number was %d", x)
	fmt.Println(jackieRobinsonReponse)

	// Printf prints a formated string to stdout
	fmt.Printf("Kobe Bryant's number was %d\n", y)
}

// SHORT DECLARATION OPERATOR (:=)
// it is shorthand for a regular variable declaration with initializer expressions but no types
// both declares a variable and assigns a value
// if we want to reassign the value to and ALREADY declared variable, we can juse use "=" instead of ":="
