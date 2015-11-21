package main

import (
	"fmt"
	"testing"
)

/*
go test -bench=.

BenchmarkEcho1 1000000	      2132 ns/op
BenchmarkEcho2 1000000	      2177 ns/op
BenchmarkEcho3 1000000	      2177 ns/op
ok  	github.com/vinceyuan/gopl-solutions/ch01/ex1.3	6.596s
*/

var (
	args = []string{"exec arg0 arg1 arg2 arg3"}
)

func TestEcho1(t *testing.T) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		echo1(args)
	} else {
		fmt.Println("no args")
	}
}

func BenchmarkEcho1(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo1(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func BenchmarkEcho2(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo2(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func BenchmarkEcho3(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo3(args)
		}
	} else {
		fmt.Println("no args")
	}
}
