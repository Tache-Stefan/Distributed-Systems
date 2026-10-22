package main

import "fmt"

func main() {
	array := []float32{3.1, 6.0, 4.8, 9.3, 10.5, 2.7}

	var index int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	res := []float32{}

	for i, num := range array {
		if i != index {
			res = append(res, num)
		}
	}

	fmt.Println(res)
}
