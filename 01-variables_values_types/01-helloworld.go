// every program must have "package main" declared
// main function is the entry point to every program

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world.")
	var t string = foo()
	fmt.Println(t)

	fizzbuzz(46)

	fmt.Println("Program over.")
}

func foo() string {
	return "I'm Terence. I'm a data engineer."
}

func fizzbuzz(i int) {
	fmt.Println("Here's some random FizzBuzz")

	for i := 0; i < 50; i++ {
		var output string = ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output != "" {
			fmt.Println(output)
		} else {
			fmt.Println(i)
		}
	}
}
