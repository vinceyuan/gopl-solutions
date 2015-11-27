package main

import (
	"bufio"
	"fmt"
	"os"
)

// Run command:
// .ex1.4 data1.txt data2.txt
func main() {
	counts := make(map[string]int)
	countsFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, countsFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsFiles map[string][]string) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !arrayContains(countsFiles[text], name) {
			countsFiles[text] = append(countsFiles[text], name)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func arrayContains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
