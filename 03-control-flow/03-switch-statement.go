package main

import (
	"fmt"
)

// switch statements perform when the evaluating statement returns true
// there is no default fall through but we can specify one (fall through keyword)
// note the fallthrough keyword will execute the next case even if it evaluates to false!
// just don't use fall through unless necessary
func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("this prints if 'fallthrough' keyword is used in the switch above")
		fallthrough
	case !(5 == 4):
		fmt.Println("we need to repeatedly specify 'fallthrough', this won't print by default")
	default:
		fmt.Println("defaults to this action of all evaluate to false")
	}

	// switch statement evaluating on a value
	n := "Bond"
	switch n {
	case "Moneypenny", "hello", "others": // we can evaluate on multiple cases in one line!
		fmt.Println("miss money")
	case "Q":
		fmt.Println("give me cool shit, Q")
	case "Bond":
		fmt.Println("bond, james")
	}
}
