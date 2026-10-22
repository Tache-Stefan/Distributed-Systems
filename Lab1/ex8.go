package main

import "fmt"

func main() {
	employees := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	departments := make(map[int][]string)

	for name, idDept := range employees {
		departments[idDept] = append(departments[idDept], name)
	}

	for id, list := range departments {
		fmt.Printf("Department %d: %v\n", id, list)
	}
}
