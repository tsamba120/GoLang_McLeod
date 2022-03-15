package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Printf("%d %s %q %t\n", x, y, y, z) // could also just use %v to get defalt value
	s := fmt.Sprintf("%d %s %q %t", x, y, y, z)
	fmt.Println(s)
}
