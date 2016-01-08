package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func count(in *bufio.Reader) (map[rune]int, [utf8.UTFMax + 1]int, int, error) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return counts, utflen, invalid, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return counts, utflen, invalid, nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	counts, utflen, invalid, err := count(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
	} else {
		fmt.Printf("rune\tcount\n")
		for c, n := range counts {
			fmt.Printf("%q\t%d\n", c, n)
		}
		fmt.Print("\nlen\tcount\n")
		for i, n := range utflen {
			if i > 0 {
				fmt.Printf("%d\t%d\n", i, n)
			}
		}
		if invalid > 0 {
			fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
		}
	}
}
