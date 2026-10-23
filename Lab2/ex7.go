package main

import "fmt"
import "sync"
import "time"

type Point struct {
	x, y int
}

func robot(id int, path []Point, out chan<- Point, in <-chan Point, wg *sync.WaitGroup) {
	defer wg.Done()

	var innerWg sync.WaitGroup
	innerWg.Add(2)

	go func() {
		defer innerWg.Done()
		for p := range in {
			fmt.Printf("[Robot %d] Received: (%d, %d)\n", id, p.x, p.y)
		}
		fmt.Printf("[Robot %d] Other robot stopped.\n", id)
	}()

	go func() {
		defer innerWg.Done()
		for _, p := range path {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("[Robot %d] Moved to: (%d, %d)\n", id, p.x, p.y)
			out <- p
		}
		close(out)
		fmt.Printf("[Robot %d] Stopped and notified the other robot.\n", id)
	}()

	innerWg.Wait()
	fmt.Printf("[Robot %d] Finished.\n", id)
}

func main() {
	ch12 := make(chan Point, 10)
	ch21 := make(chan Point, 10)

	var wg sync.WaitGroup

	path1 := []Point{{0, 1}, {1, 1}, {2, 1}, {2, 2}}
	path2 := []Point{{10, 10}, {10, 11}, {11, 11}, {12, 12}, {13, 12}}

	wg.Add(2)
	go robot(1, path1, ch12, ch21, &wg)
	go robot(2, path2, ch21, ch12, &wg)

	wg.Wait()
	fmt.Println("[Observer] Both robots finished.")
}
