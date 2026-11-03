package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done bool
}

func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range in {
		if msg.Done {
			if (!*isLeader) {
				out <- msg
			}
			return
		}

		if (msg.Value > id) {
			out <- msg
		} else if (msg.Value == id) {
			*isLeader = true
			out <- Msg{id, true}
		}
	}
}

func RingFlooding(ids []int) int {
	chans := make([]chan Msg, len(ids))
	for i := range chans {
		chans[i] = make(chan Msg, 1)
	}
	is_leader := make([]bool, len(ids))
	var wg sync.WaitGroup

	for i, val := range ids {
		wg.Add(1)
		next_idx := (i + 1) % len(ids)
		chans[next_idx] <- Msg{val, false}
		go node(val, chans[i], chans[next_idx], &is_leader[i], &wg)
	}

	wg.Wait()

	for i, val := range is_leader {
		if (val == true) {
			return ids[i]
		}
	}

	return -1
}

func main() {
	ids := []int{7, 4, 9, 2, 5, 33, 10, 1}
	leader := RingFlooding(ids)
	fmt.Printf("Leader: %d\n", leader)
}
