package main

import "fmt"

type car struct {
	name  string
	power int
}

type supercar struct {
	car   car
	color string
}

func main() {
	c1 := supercar{
		car: car{
			name:  "Ferrari",
			power: 800,
		},
		color: "blue",
	}

	fmt.Println(c1)
}
