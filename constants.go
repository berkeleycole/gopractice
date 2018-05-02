package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// Go supports constants of character, string, boolean, and numeric values.

func main() {

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}