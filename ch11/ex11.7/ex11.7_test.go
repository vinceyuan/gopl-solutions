package intset

import (
	"math/rand"
	"testing"
	"time"
)

func benchmarkAdd(b *testing.B, n int) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		x.Add(n)
	}
}

func BenchmarkAdd1(b *testing.B) {
	benchmarkAdd(b, 1)
}
func BenchmarkAdd10(b *testing.B) {
	benchmarkAdd(b, 10)
}
func BenchmarkAdd100(b *testing.B) {
	benchmarkAdd(b, 100)
}
func BenchmarkAdd1000(b *testing.B) {
	benchmarkAdd(b, 1000)
}

func BenchmarkAddPseudoRandom(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	var x IntSet
	for i := 0; i < b.N; i++ {
		n := rng.Intn(1000000)
		x.Add(n)
	}
}
