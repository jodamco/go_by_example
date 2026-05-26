package main

import "fmt"

func main() {

	// Short declaration statement
	// Always inferred
	i := 1

	// Single condition for
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial/condition/after loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Basic 'do this n times' - this is probably looping through the items of an array or so
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Conditionless loop - can be broken out with 'break'
	for {
		fmt.Println("loop")
		break
	}

	// It is possible to skip iterations with 'continue'
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
