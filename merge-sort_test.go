package mergesort

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSequentialMergesort(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	size := 10000
	s := make([]int, size)
	for i := range s {
		s[i] = rand.Intn(size)
	}

	for i := 0; i < b.N; i++ {
		sCopy := make([]int, len(s))
		copy(sCopy, s)

		sequentialMergesort(sCopy)
	}
}

func BenchmarkParallelMergesortV1(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	size := 10000
	s := make([]int, size)
	for i := range s {
		s[i] = rand.Intn(size)
	}

	for i := 0; i < b.N; i++ {
		sCopy := make([]int, len(s))
		copy(sCopy, s)

		parallelMergesortV1(sCopy)
	}
}

func BenchmarkParallelMergesortV2(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	size := 10000
	s := make([]int, size)
	for i := range s {
		s[i] = rand.Intn(size)
	}

	for i := 0; i < b.N; i++ {
		sCopy := make([]int, len(s))
		copy(sCopy, s)

		parallelMergesortV2(sCopy)
	}
}
