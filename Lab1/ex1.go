package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for i := 2; i <= num / 2; i++ {
		if num % i == 0 {
			fmt.Println(num, "is not prime.")
			return
		}
	}

	fmt.Println(num, "is prime.")
}
