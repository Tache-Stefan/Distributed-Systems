package main

import "fmt"

func send(ch chan string) {
	ch <- "msg"
}

func recv(ch chan string, chOut chan string) {
	msg := <- ch
	chOut <- msg
}

func main() {
	ch := make(chan string)
	chOut := make(chan string)

	go send(ch)
	go recv(ch, chOut)

	fmt.Println(<- chOut)
}
