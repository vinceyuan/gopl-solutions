// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string...
package main

import (
	"unicode/utf8"
	"fmt"
)

func reverseBytes(a []byte)[]byte {
	if utf8.RuneCount(a) == 1 {
		return a
	}
	_,s := utf8.DecodeRune(a)
	return append(reverseBytes(a[s:]), a[:s]...)
}

func main() {
	a := []byte("abb")
	fmt.Println(string(reverseBytes(a)))
}
