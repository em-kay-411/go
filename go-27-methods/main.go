package main

import "fmt"

type animals struct {
	name string
}

func (a animals) display() {
	fmt.Println(a)
}

func main() {
	animal := animals{
		name: "dog",
	}

	animal.display()
}
