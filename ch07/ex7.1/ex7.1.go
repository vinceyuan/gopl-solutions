package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	c := 0
	for scanner.Scan() {
		c++
	}
	*wc += WordCounter(c)
	return c, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	c := 1
	for _, b := range p {
		if b == '\n' {
			c++
		}
	}
	*lc += LineCounter(c)
	return c, nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("Hello word! 你好"))
	fmt.Println(wc)
	var lc LineCounter
	lc.Write([]byte(`hello
abc
def ad fas 
`))
	fmt.Println(lc)
}
