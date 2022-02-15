package main

import "fmt"

func main() {
	// Declare variable using var keyword
	var x int
	var y int
	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean int
	mean = (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)

	// Declare variable using :=
	a := 1.0
	b := 2.0

	fmt.Printf("x=%v, type of %T\n", a, a)
	fmt.Printf("y=%v, type of %T\n", b, b)

	average := (a + b) / 2.0
	fmt.Printf("result: %v, type of %T\n", average, average)

	// Declare 2 variables in single-line
	i, j := 1, 2
	fmt.Printf("i=%v, and j=%v, type of i=%T, and type of j=%T", i, j, i, j)
}
