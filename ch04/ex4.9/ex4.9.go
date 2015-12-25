package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode wrod

	fileReader, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\t\tfreq\n")
	for c, n := range counts {
		fmt.Printf("%q\t\t%d\n", c, n)
	}
}
