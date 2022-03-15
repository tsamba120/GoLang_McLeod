package main

import "fmt"

// From documentation:
// Package fmt implements formatted I/O with functions analogous to  C's printf and scanf
// The format 'verbs' are derived from C's but are simpler
// See documentation on info on verbs (i.e. %T = type, %s = string, $d = digit)

var y int = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // print type of var
	fmt.Printf("%b\n", y)  // print as binary
	fmt.Printf("%#x\n", y) // print as hexidemical
	fmt.Printf("%#x\n", y) // print as hexidemical with Go-syntax "0x" in front
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

}

// 3 important func groups
// 1. general printing - Print, Printf Println, etc.
// 2. Printing to a string that can be assigned - Sprint(), Sprintf(), Sprintln()
// 3. Printing to a file or a web server's response - Fprint(), Fprintf(), Fprintln()

// %v - the value in the default format
// %#v - a Go-syntax representation of the value
// %T - a Go-syntax representation of the type of the value
// %% a literal percent sign; consumes no value

// other verbs exist for strings, ints, etc.
