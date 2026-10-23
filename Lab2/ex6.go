package main

import "fmt"

func start(ch chan bool) {
	ch <- true
}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go start(ch1)
	go start(ch2)

	select {
	case <- ch1:
		fmt.Println("Robot 1 finished first")
	case <- ch2:
		fmt.Println("Robot 2 finished first")
	}
}
