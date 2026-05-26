package main

import (
	"fmt"
	"slices"
)

func main() {

	// An uninitialized slice equals to nil and has len 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create a slice with non-zero length we use make
	// Make should be used since slices need initialization of the underlying structure
	// By default, an slice capacity is equal to its length
	s = make([]string, 3)
	fmt.Println("emp:", s, "length:", len(s), "cap:", cap(s))

	// Set is made in the same way as with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Slices support append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy'd
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Slices support slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(s)
	whatAmI(l)
	var r [3]string
	whatAmI(r)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
