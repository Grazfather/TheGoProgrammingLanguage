package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkPopCountTable(b *testing.B) {
	rand.Seed(int64(5))
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(rand.Int63()))
	}
}
func BenchmarkPopCountLoop(b *testing.B) {
	rand.Seed(int64(5))
	for i := 0; i < b.N; i++ {
		PopCountLoop(uint64(rand.Int63()))
	}
}
func BenchmarkPopCountShift(b *testing.B) {
	rand.Seed(int64(5))
	for i := 0; i < b.N; i++ {
		PopCountShift(uint64(rand.Int63()))
	}
}
func BenchmarkPopCountClear(b *testing.B) {
	rand.Seed(int64(5))
	for i := 0; i < b.N; i++ {
		PopCountClear(uint64(rand.Int63()))
	}
}
