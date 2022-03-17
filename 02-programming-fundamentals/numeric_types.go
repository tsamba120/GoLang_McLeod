// see course materials pdf for numeric types

// uint - byte - uints are "unsigned" integers. 0 and above
// int - rune
// floats

// we can adjust the size of these variables (i.e. int32, uint64) all base-2 digits

// package runtime

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   // print operating system - darwin
	fmt.Println(runtime.GOARCH) // print architecture - arm64
}
