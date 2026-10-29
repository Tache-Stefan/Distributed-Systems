package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2

	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	s_clock := Clock{offset: 0}

	initialOffset := -200 * time.Millisecond
	c_clock := Clock{offset: initialOffset}

	fmt.Println("Single measurement:")
	thetaA, deltaA := measure(c_clock, s_clock)
	fmt.Printf("Initial offset:     %v\n", initialOffset)
	fmt.Printf("Estimated theta:    %v\n", thetaA)
	fmt.Printf("Delta (Round-Trip): %v\n", deltaA)

	adjustedOffsetA := initialOffset + thetaA
	fmt.Printf("Remaining client offset: %v\n", adjustedOffsetA)

	fmt.Println("\n10 measurements:")
	c_clock = Clock{offset: initialOffset}

	var bestTheta, minDelta time.Duration

	for i := 1; i <= 10; i++ {
		theta, delta := measure(c_clock, s_clock)
		fmt.Printf("Measurement %d: theta=%v, delta=%v\n", i, theta, delta)

		if i == 1 || delta < minDelta {
			bestTheta = theta
			minDelta = delta
		}
	}

	fmt.Printf("\nBest theta: %v with minimum delta: %v\n", bestTheta, minDelta)
	adjustedOffsetB := initialOffset + bestTheta
	fmt.Printf("Remaining client offset after best measurement: %v\n", adjustedOffsetB)
}
