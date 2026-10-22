package main

import "fmt"

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	var l, r int = 0, len(slice) - 1
	for l < r {
		slice[l], slice[r] = slice[r], slice[l]
		l++
		r--
	}

	fmt.Println(slice)
}
