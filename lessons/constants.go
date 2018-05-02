package lessons

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean, and numeric values.
const s string = "constant"

// Constants ...
func Constants() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}
