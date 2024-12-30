package main

import "fmt"

func main() {
	map1 := map[string]int{"John": 23, "Robin": 20}
	fmt.Println(map1)

	map2 := make(map[string]int)
	map2["John"] = 23
	map2["Robin"] = 20
	fmt.Println(len(map2))

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
