package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n, &sum)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	seen := make(map[int]int)

	for i, x := range v {
		complement := sum - x
		idx, ok := seen[complement]
		if ok {
			fmt.Println(idx, i)
		}
		seen[x] = i
	}
}
