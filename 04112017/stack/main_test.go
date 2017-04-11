package main

import (
	"testing"
)

func BenchmarkStackMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stackMap()
	}
}

func BenchmarkHeapMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heapMap()
	}
}
