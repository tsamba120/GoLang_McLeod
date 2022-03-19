package main

import (
	"fmt"
)

// Write a program that prints a number in decimal, binary, and hex

var x int = 56

func main() {
	fmt.Printf("%d\t\t%b\t\t%x\n", x, x, x)
}
