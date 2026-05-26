package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// A typeless constant
	const n = 500000000

	const d float32 = 3e20 / n
	fmt.Println(d)

	// Converts d into an int
	fmt.Println(int64(d))

	// Converts n into a float since math.Sin expects that type
	fmt.Println(math.Sin(n))
}
