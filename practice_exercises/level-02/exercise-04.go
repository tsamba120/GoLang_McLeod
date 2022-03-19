package main

import (
	"fmt"
)

// Write a program that assigns an int variable
// Print in decimal, binary, hex
// Shift bits 1 position to the left, assign to new variable
// Print in decimal, binary, hex once again

var x int = 64

func main() {
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %x\n\n", x)

	fmt.Println("Bit shifting one position to the left...")

	y := x << 1
	fmt.Printf("\nDecimal: %d\n", y)
	fmt.Printf("Binary: %b\n", y)
	fmt.Printf("Hex: %x\n", y)

}
