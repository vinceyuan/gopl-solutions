package popcount

/*
$ go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkPopCount-4          	200000000	         8.00 ns/op
BenchmarkPopCount2-4         	100000000	        16.3 ns/op
BenchmarkPopCountByShifting-4	20000000	        96.8 ns/op
ok  	github.com/vinceyuan/gopl-solutions/ch02/ex2.4/popcount	6.197s
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

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}
