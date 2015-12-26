package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("hello,$foo", toAbc))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func toAbc(s string) string {
	return "abc"
}
