package main

import (
	"fmt"
	"time"
)

// Not tested, because it never runs out of memory on my Mac.
func f() {
	ch := make(chan struct{})
	begin := time.Now()
	var count int64 = 0

	defer func(begin time.Time) {
		end := time.Now()
		fmt.Printf("\nDone after %d seconds\n", end.Sub(begin)/1000000)
		if p := recover(); p != nil {
			fmt.Print("Out of memory. Sending data to all goroutines...\n")
			begin = time.Now()
			for ; count > 0; count-- {
				ch <- struct{}{}
			}
			end = time.Now()
			fmt.Printf("Done after %d seconds\n", end.Sub(begin)/1000000)
		}
	}(begin)

	fmt.Println("Number of goroutines:")
	for {
		count++
		fmt.Printf("\r%d", count)
		go func() { <-ch }()
	}
}

func main() {
	f()
}
