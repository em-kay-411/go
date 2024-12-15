package main

import "fmt"

var name string = "rohan"

func main() {
	var a int = 54
	// var name string = "rohan"
	b := 4 // Short declarration operator
	fmt.Printf("%b %x \n", a, a)
	fmt.Println(a)
	fmt.Println("hello world", a, "go")
	fmt.Println(b)
	fmt.Printf("%q", name)
}
