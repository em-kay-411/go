package main

import "fmt"

func main() {
	map1 := map[string]int{"John": 22, "Lucy": 45}
	fmt.Println(map1["John"])
	map1["Rohan"] = 35
	fmt.Println(map1)

	map2 := make(map[string]int)
	map2["John"] = 22
	map2["Lucy"] = 45

	fmt.Println(map2)
}
