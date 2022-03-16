package main

import "fmt"

type kitten int

var x kitten

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
