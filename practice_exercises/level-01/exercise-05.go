package main

import "fmt"

type kitten int

var x kitten
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 562
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
