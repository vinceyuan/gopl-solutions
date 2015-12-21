package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	sign := ""
	if s[:1] == "+" || s[:1] == "-" {
		sign = s[:1]
		s = s[1:]
	}

	hasDot := false
	n := len(s)
	var nBeforeDot int
	if indexOfDot := strings.Index(s, "."); indexOfDot >= 0 {
		hasDot = true
		nBeforeDot = indexOfDot
	} else {
		nBeforeDot = n
	}

	// Handle digits before dot
	i := nBeforeDot % 3
	if i == 0 {
		i += 3
	}
	for i < nBeforeDot {
		s = s[:i] + "," + s[i:]
		i += 4
		nBeforeDot++
		n++
	}

	// Handle digits after dot
	if hasDot {
		i = nBeforeDot + 1 + 3
		for i < n {
			s = s[:i] + "," + s[i:]
			i += 4
			n++
		}
	}

	// Add sign
	if sign != "" {
		s = sign + s
	}
	return s
}
