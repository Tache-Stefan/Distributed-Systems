package main

import "fmt"

func send(ch chan string) {
	ch <- "salut"
}

func main() {
	ch := make(chan string)
	go send(ch)
	fmt.Println(<-ch)
}
