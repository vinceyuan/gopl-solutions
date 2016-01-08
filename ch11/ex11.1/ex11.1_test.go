package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		char  rune
		count int
	}{
		{"abcd", 'a', 1},
		{"⌘⌘⌘⌘⌘", '⌘', 5},
		{"hello world\n hello", 'l', 5},
		{"hello world\n hello\n", '\n', 2},
		{"你好hello你好", '好', 2},
	}

	for _, test := range tests {
		in := bufio.NewReader(strings.NewReader(test.input))
		counts, _, _, _ := count(in)
		if counts[test.char] != test.count {
			t.Errorf("counts of %v = %v", test.char, counts[test.char])
		}
	}

}
