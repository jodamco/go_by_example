package main

import (
	"fmt"
	"maps"
)

func main() {

	// Maps also requires make
	// map[key-type]val-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("len:", len(m))

	delete(m, "k45")
	fmt.Println("len:", len(m))

	clear(m)
	fmt.Println("map:", m)

	// The map value getter returns 2 values:
	// 1st - the value itself/ 2nd - a bool indicating if the value was present
	// this is used to disambiguate between missing keys and keys with zero values
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
