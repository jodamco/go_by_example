package main

import "fmt"

// Variadic function. It can be called with any number of trailing args.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// Range is a function that iterates through an array, slice, map, etc
	// It returns the index and the item in this particular case
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{
		1, 2, 3, 4,
	}

	// Adding ... after the slice, spreads it.
	sum(nums...)
}
