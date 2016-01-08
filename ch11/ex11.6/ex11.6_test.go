package popcount

/*
$ go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkPopCount-4          	200000000	         7.84 ns/op
BenchmarkPopCount2-4         	100000000	        16.2 ns/op
BenchmarkPopCountByShifting-4	20000000	        96.9 ns/op
BenchmarkPopCountByClearing-4	30000000	        38.6 ns/op
ok  	github.com/vinceyuan/gopl-solutions/ch02/ex2.4/popcount	7.322s
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

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}
