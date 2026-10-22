package main

import "fmt"

func main() {
	name_height := map[string]int{
		"Andrei": 190,
		"Maria": 165,
		"Ion": 182,
		"Elena": 172,
		"George": 178,
		"Ana": 165,
	}

	names := []string{}
	for name := range name_height {
		names = append(names, name)
	}

	for i := 0; i < len(names) - 1; i++ {
		for j := i + 1; j < len(names); j++ {
			if name_height[names[i]] < name_height[names[j]] {
				names[i], names[j] = names[j], names[i]
			}
		}
	}

	fmt.Println(names)
}
