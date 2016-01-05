package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	ch3 := make(chan struct{})

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		i := 0
		seconds := 0
		fmt.Println("Number of communications per second:")
	loop:
		for {
			ch1 <- struct{}{}
			select {
			case <-ch2:
				i++
			case <-ticker.C:
				seconds++
				<-ch2 // Must consume ch2, otherwise dead lock
				i++
				fmt.Printf("\rSecond: %d: %d", seconds, 2*i/seconds)
				if seconds >= 5 {
					ticker.Stop()
					break loop
				}
			}
		}

		ch3 <- struct{}{}
	}()

	go func() {
		for {
			<-ch1
			ch2 <- struct{}{}
		}
	}()

	<-ch3
	fmt.Println()
}
