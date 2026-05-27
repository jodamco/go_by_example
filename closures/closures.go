package main

import "fmt"

// Defines a function that returns a function
func intSeq() func() int {
	i := 0

	// Anonymous fn declaration
	return func() int {
		// This fn enclosure i, thus creating
		// some sort of 'stateful' fn
		i++
		return i
	}
}

func main() {

	// Creates the closure
	nextInt := intSeq()

	// i = 0
	fmt.Println(nextInt())
	// i = 1
	fmt.Println(nextInt())
	// i = 2
	fmt.Println(nextInt())
	// i = 3

	// Separate closure
	newInts := intSeq()

	// newI = 0
	fmt.Println(newInts())
	// newI = 1
}
