package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	
	err = scanner.Err()
	if err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		return
	}

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
