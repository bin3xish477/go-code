package main

import (
	"fmt"
)

// ExportedConst is a constant that will can be exported
const ExportedConst = 1000

const (
	e = iota
	f = iota
	g = iota
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// All primitive type constants:
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true

	fmt.Println("Printing all constant primitive types...")
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	fmt.Println("Printing iota constants...")
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
}
