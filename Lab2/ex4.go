package main

import "fmt"

func process(ch chan bool) {
	fmt.Println("Processing...")
	fmt.Println("Process done")
	ch <- true
}

func main() {
	ch := make(chan bool)

	go process(ch)

	<- ch
	fmt.Println("Main done")
}
