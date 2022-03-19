package main

import (
	"fmt"
)

// %d is digits
// %b is binary value of integer

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Println("Shift bits to the left by 1 unit")
	fmt.Printf("%d\t\t%b\n", y, y)

	z := y << 2
	fmt.Println("Shift bits to the left by 2 units")
	fmt.Printf("%d\t\t%b\n", z, z)

	var a byte = byte('h')
	fmt.Println("\nRight shift example: ")
	fmt.Printf("%b\n", a)

	b := a >> 1
	fmt.Printf("%b\n\n", b)

	// Let's look at kilobyte, megabyte, and gigabytes

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	// Note how it shifts to the left each time
	// fmt.Printf("%d\t\t\t%b\n", kb, kb)
	// fmt.Printf("%d\t\t\t%b\n", mb, mb)
	// fmt.Printf("%d\t\t%b\n", gb, gb)

	iotaExample()

}

// We'll use iota to accomplish the same thing

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func iotaExample() {
	fmt.Println("Using iota for bit-shifting: ")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
