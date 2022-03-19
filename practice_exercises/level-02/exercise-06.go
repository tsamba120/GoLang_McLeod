package main

import (
	"fmt"
)

// Using iota, create 4 constants for the next 4 years. Print the constant values

const (
	year1 = 2022
	year2 = year1 + iota
	year3 = year1 + iota
	year4 = year1 + iota
)

func main() {
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

}
