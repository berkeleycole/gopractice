package lessons

import "fmt"

// ValuesAndVariables ...
func ValuesAndVariables() {
	// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

	// TEXT
	fmt.Println("go" + "lang")

	// BOOLEANS
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// INTEGERS
	fmt.Println("Math with Integers: 1+1 =", 1+1)

	// FLOATS
	fmt.Println("Math with Floats: 7.0/3.0 =", 7.0/3.0)

	// VARIABLES
	var a = "initial"
	fmt.Println("variables: " + a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
	f := "short"
	fmt.Println(f)
}
