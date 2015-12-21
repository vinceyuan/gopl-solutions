package main

import (
	"fmt"
	"unicode"
)

func removeDupSpace(b []byte) []byte {
	out := b[:0]
	for i, c := range b {
		if unicode.IsSpace(rune(c)) {
			if i > 0 && unicode.IsSpace(rune(b[i-1])) {
				continue
			} else {
				out = append(out, ' ')
			}
		} else {
			out = append(out, c)
		}
	}
	return out
}

func main() {
	b := []byte("abc\r  \n\rdef")
	fmt.Printf("%q\n", string(removeDupSpace(b)))
	fmt.Printf("%q\n", b)
}
