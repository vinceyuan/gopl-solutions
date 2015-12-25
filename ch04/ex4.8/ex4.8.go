package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	CHAR_IS_SPACE = iota
	CHAR_IS_SYMBOL
	CHAR_IS_MARK
	CHAR_IS_DIGIT
	CHAR_IS_PRINT
	CHAR_IS_PUNCT
	CHAR_IS_LETTER
	CHAR_IS_NUMBER
	CHAR_IS_CONTROL
	CHAR_IS_GRAPHIC
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsSpace(r):
			counts[CHAR_IS_SPACE]++
		case unicode.IsSymbol(r):
			counts[CHAR_IS_SYMBOL]++
		case unicode.IsMark(r):
			counts[CHAR_IS_MARK]++
		case unicode.IsDigit(r):
			counts[CHAR_IS_DIGIT]++
		case unicode.IsPrint(r):
			counts[CHAR_IS_PRINT]++
		case unicode.IsPunct(r):
			counts[CHAR_IS_PUNCT]++
		case unicode.IsLetter(r):
			counts[CHAR_IS_LETTER]++
		case unicode.IsNumber(r):
			counts[CHAR_IS_NUMBER]++
		case unicode.IsControl(r):
			counts[CHAR_IS_CONTROL]++
		case unicode.IsGraphic(r):
			counts[CHAR_IS_GRAPHIC]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	var tname string
	for t, n := range counts {
		switch t {
		case CHAR_IS_SPACE:
			tname = "space"
		case CHAR_IS_SYMBOL:
			tname = "symbol"
		case CHAR_IS_MARK:
			tname = "mark"
		case CHAR_IS_DIGIT:
			tname = "digit"
		case CHAR_IS_PRINT:
			tname = "print"
		case CHAR_IS_PUNCT:
			tname = "punct"
		case CHAR_IS_LETTER:
			tname = "letter"
		case CHAR_IS_NUMBER:
			tname = "number"
		case CHAR_IS_CONTROL:
			tname = "control"
		case CHAR_IS_GRAPHIC:
			tname = "graphic"
		}

		fmt.Printf("%s\t%d\n", tname, n)
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
