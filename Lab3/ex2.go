package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

func less(a, b Event) bool {
	if a.internal_counter == b.internal_counter {
		return a.process < b.process
	}
	return a.internal_counter < b.internal_counter
}

func main() {
	const numProcesses = 3
	var wg sync.WaitGroup

	channels := make([]chan Message, numProcesses)
	for i := 0; i < numProcesses; i++ {
		channels[i] = make(chan Message, 10)
	}

	eventChan := make(chan Event, 50)

	for i := 0; i < numProcesses; i++ {
		wg.Add(1)
		pid := i + 1
		inCh := channels[(i + numProcesses - 1) % numProcesses]
		outCh := channels[i]

		go func(id int, in <-chan Message, out chan<- Message) {
			defer wg.Done()
			var clock LamportClock

			internalTicks := rand.Intn(3) + 1
			for t := 0; t < internalTicks; t++ {
				clock.Tick()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "internal"}
			}

			if id == 1 {
				ts := clock.Send()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "send"}
				out <- Message{sender: id, ts: ts}

				msg := <-in
				clock.Receive(msg.ts)
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "recv"}
			} else {
				msg := <-in
				clock.Receive(msg.ts)
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "recv"}

				clock.Tick()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "internal"}

				ts := clock.Send()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "send"}
				out <- Message{sender: id, ts: ts}
			}
		}(pid, inCh, outCh)
	}

	wg.Wait()
	close(eventChan)

	var events []Event
	for ev := range eventChan {
		events = append(events, ev)
	}

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	for _, ev := range events {
		fmt.Printf("P%-9d | %-10d | %-15s\n", ev.process, ev.internal_counter, ev.kind)
	}
}
