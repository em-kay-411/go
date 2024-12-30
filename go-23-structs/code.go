package main

import "fmt"

type car struct {
	name  string
	power int
}

func main() {
	c1 := car{
		name:  "Ferrari",
		power: 800,
	}

	fmt.Println(c1)
}
