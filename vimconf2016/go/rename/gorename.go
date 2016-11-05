// https://gobyexample.com/for
// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// The most basic type, with a single condition.
	wtfnum := 1
	for wtfnum <= 3 {
		fmt.Println(wtfnum)
		wtfnum = wtfnum + 1
	}

	// A classic initial/condition/after `for` loop.
	for i := 7; i <= 9; i++ {
		fmt.Println(i)
		log.Printf("test")
		_ = os.Getpid()
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}
}
