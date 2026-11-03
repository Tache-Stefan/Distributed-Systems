package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value float64
}

func node(id int, initialVote float64, in []<-chan Msg, out []chan<- Msg, rounds int, finalAvg *float64, wg *sync.WaitGroup) {
	defer wg.Done()

	currentVal := initialVote
	dIn := float64(len(in))

	for r := 0; r < rounds; r++ {
		for _, ch := range out {
			ch <- Msg{Value: currentVal}
		}

		sumNeighbors := 0.0
		for _, ch := range in {
			msg := <-ch
			sumNeighbors += msg.Value
		}

		currentVal = (currentVal + sumNeighbors) / (dIn + 1.0)
	}

	*finalAvg = currentVal
}

func MajorityVoting(nodes []int, initialVotes map[int]float64, adj map[int][]int, rounds int) (bool, float64) {
	var wg sync.WaitGroup

	finalAverages := make([]float64, len(nodes))
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
		go node(id, initialVotes[id], inChans[id], outChans[id], rounds, &finalAverages[i], &wg)
	}

	wg.Wait()

	fmt.Println("Local results")
	for i, id := range nodes {
		fmt.Printf("  Node %2d has converged to value: %.6f\n", id, finalAverages[i])
	}

	consensusAvg := finalAverages[0]
	majorityDecision := consensusAvg > 0.5

	return majorityDecision, consensusAvg
}

func main() {
	nodes := []int{14, 19, 22, 8}

	adj := map[int][]int{
		14: {19, 22},
		19: {22, 8},
		22: {14, 8},
		8:  {14, 19},
	}

	votes := map[int]float64{
		14: 1.0,
		19: 1.0,
		22: 0.0,
		8:  1.0,
	}

	fmt.Println("Initial votes")
	for id, v := range votes {
		fmt.Printf("  Node %d: %.0f\n", id, v)
	}

	rounds := 20

	decision, avg := MajorityVoting(nodes, votes, adj, rounds)

	fmt.Printf("Median: %.4f\n", avg)
	if decision {
		fmt.Println("Majority 1")
	} else {
		fmt.Println("Majority 0")
	}
}
