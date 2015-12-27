package main

import (
	"fmt"
	"math"
)

func max(vals ...int) (result int, ok bool) {
	result = math.MinInt64
	if len(vals) == 0 {
		return
	}
	ok = true
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return
}

func min(vals ...int) (result int, ok bool) {
	result = math.MaxInt64
	if len(vals) == 0 {
		return
	}
	ok = true
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
}
