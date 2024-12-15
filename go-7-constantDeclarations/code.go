package main

import "fmt"

func main() {
	const a int = 1 // typed
	const b = 2     // untyped

	const (
		d     = 2.0
		f     = 24.56
		r     = "hello"
		iota1 = iota
		iota2 = iota
		iota3 = iota
	)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("d", d)
	fmt.Println("f", f)
	fmt.Println("r", r)
	fmt.Println("iota1", iota1)
	fmt.Println("iota2", iota2)
	fmt.Println("iota3", iota3)
}
