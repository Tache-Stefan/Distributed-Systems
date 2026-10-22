package main

import "fmt"

func main() {
	array := []int{3, 6, 4, 9, 10, 2}

	for i := 0; i < len(array) - 1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	fmt.Println(array)
}
