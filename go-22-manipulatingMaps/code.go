package main

import "fmt"

func main() {
	map1 := map[string]int{"John": 22, "Lucy": 45}

	// Adding element
	map1["Rohan"] = 35
	fmt.Println(map1)

	// Deleting element
	delete(map1, "John")
	fmt.Println(map1)

	// Search in map
	if _, ok := map1["John"]; !ok {
		fmt.Println("John is not present")
	} else {
		fmt.Println("John is present")
	}

	// Iterating
	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
