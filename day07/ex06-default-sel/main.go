package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// creates a channel that sends a signal every 100 milliseconds.
	tick := time.Tick(100 * time.Millisecond)
	// creates a channel that sends a signal only once, after 500 milliseconds.
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
