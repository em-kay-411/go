package main

import "fmt"

func main() {
	map1 := map[string]int{"John": 23, "Robin": 20}
	fmt.Println(map1)

	// Add new element
	map1["Rohan"] = 34

	// Delete an element
	delete(map1, "John")

	// Check if the key is present in a map
	if value, ok := map1["John"]; !ok {
		fmt.Println("The key is not present in the map")
	} else {
		fmt.Println(value)
	}
}
