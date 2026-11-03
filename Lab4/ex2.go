package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
}

func node(id int, in []<-chan Msg, out []chan<- Msg, rounds int, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	maxID := id

	for r := 0; r < rounds; r++ {
		for _, ch := range out {
			ch <- Msg{Value: maxID}
		}

		for _, ch := range in {
			msg := <-ch
			if msg.Value > maxID {
				maxID = msg.Value
			}
		}
	}

	if maxID == id {
		*isLeader = true
	}
}

func GraphFlooding(nodes []int, adj map[int][]int, diameter int) int {
	var wg sync.WaitGroup

	isLeader := make([]bool, len(nodes))
	inChans := make(map[int][]<-chan Msg)
	outChans := make(map[int][]chan<- Msg)

	for u, neighbors := range adj {
		for _, v := range neighbors {
			ch := make(chan Msg, 1)
			outChans[u] = append(outChans[u], ch)
			inChans[v] = append(inChans[v], ch)
		}
	}

	for i, id := range nodes {
		wg.Add(1)
		go node(id, inChans[id], outChans[id], diameter, &isLeader[i], &wg)
	}

	wg.Wait()

	for i, leader := range isLeader {
		if leader {
			return nodes[i]
		}
	}

	return -1
}

func main() {
	nodes := []int{14, 19, 22, 8}

	adj := map[int][]int{
		14: {19, 22},
		19: {22, 8},
		22: {14, 8},
		8:  {14, 19},
	}

	diameter := 2

	leader := GraphFlooding(nodes, adj, diameter)
	fmt.Printf("Leader: %d\n", leader)
}
