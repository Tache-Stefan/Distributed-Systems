package main

import "fmt"

func send(ch chan string, msg string) {
	ch <- msg
}

func main() {
	ch := make(chan string)

	go send(ch, "msg1")
	go send(ch, "msg2")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
