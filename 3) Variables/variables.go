package main

import "fmt"

func main() {
	var a = " test "
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// You can create a variable without the initial value
	// Go will assign a default value ,for integer variables "Zero" , and for
	// String variables , Empty string ("") , and for boolean variables "False"
	var e int
	fmt.Println("Empty integer variable : ", e)

	var g bool
	fmt.Println("Empty boolean variable : ", g)

	/*
		The := syntax is shorthand for declaring and initializing a variable,
		e.g. for var f string = "short" in this case.
	*/

	f := "short"
	fmt.Println(f)
}
