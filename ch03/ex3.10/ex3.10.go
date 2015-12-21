package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	i := n % 3
	if i == 0 {
		i += 3
	}
	for i < n {
		s = s[:i] + "," + s[i:]
		i += 4
		n++
	}
	return s
}
