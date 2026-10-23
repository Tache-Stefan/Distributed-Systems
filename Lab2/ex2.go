package main

import "fmt"
import "sync"

func recv(ch chan string, wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	go recv(ch, &wg)
	wg.Add(1)

	ch <- "salut"
	wg.Wait()
}
