package popcount

/*
$ go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkPopCount-4 	200000000	         7.93 ns/op
BenchmarkPopCount2-4	100000000	        16.3 ns/op
ok  	github.com/vinceyuan/gopl-solutions/ch02/ex2.3/popcount	4.240s
*/

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(0x1234567890ABCDEF)
	}
}
