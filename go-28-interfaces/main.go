package main

import "fmt"

type Dog struct {
	limbs int
	age   int
	breed string
}

type Cat struct {
	limbs int
	age   int
}

type Animal interface {
	Sound()
}

func (d Dog) Sound() {
	fmt.Println("bark")
}

func (c Cat) Sound() {
	fmt.Println("meow")
}

func makeSound(a Animal) {
	a.Sound()
}

func main() {
	d := Dog{
		limbs: 4,
		age:   5,
		breed: "husky",
	}

	c := Cat{
		limbs: 4,
		age:   3,
	}

	makeSound(d)
	makeSound(c)
}
