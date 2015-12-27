package main

import "fmt"

func Join(sep string, a ...string) string {
	var out string
	length := len(a)
	for i, s := range a {
		out = out + s
		if i != length-1 {
			out = out + sep
		}
	}
	return out
}

func main() {
	fmt.Println(Join(", ", "aaa", "bbb", "ccc"))
}
