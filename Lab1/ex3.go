package main

import "fmt"

func main() {
	array := []int{3, 6, 4, 9, 10, 2}
	res := []int{}

	for _, num := range array {
		if num % 2 == 0 {
			res = append(res, num)
		}
	}

	fmt.Println(res)
}
