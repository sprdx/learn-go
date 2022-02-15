// To refresh my knowledge about Go, I'm here to code basic of Go from Go Essential Training origin from Linkedin course.

// So, a comment in Go could be written using double slash for inline comment

/*
and slash followed by asterisk and asterisk then slash for multi-line comment
*/

// In Go, code is organized in packages
// Package main has special function for compiling the code to be executable file
package main

// Import statement is used to source from other package or library
import (
	"fmt"
)

// Function main is used to be executed by Go Runtime
func main() {
	// function is invoked with prefix of package that function is came from
	fmt.Println("Hello, World!!")
}
