package main

import (
	"fmt"
)
/* 
Variables that are capitilized are exported.
Variables that are internal, are lowered case. 
*/

// package level const
const a uint16 = 27

func main() {
	fmt.Printf("Global variable a = %d\n", a)
	
	f := 3.14
	fmt.Printf("Pie = %f\n", f)

	s := "This is a string" // implicitly declaration of string
	fmt.Printf("s = %s\n", s) // no string formating, print new line

	b := []byte(s) // byte array of string s, in others word utf-8 array of the string argument
	fmt.Printf("This is a byte array of the string s:\n%v\n", b) // formated I/O

	// - runes are int32
	// - runes are not commonly used
	r := 'a' // rune data type
	fmt.Printf("Char %c = %v in and %b in binary\n", r, r, r) // prints the utf-8 equivalent of rune 'a'

	/* Constants in Golang */
	const myCount int = 42 // declaring constants
	fmt.Printf("This is a constant: %v\n", myCount)
}
