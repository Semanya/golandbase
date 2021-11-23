package main

import (
	// "fmt"
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

var pool = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func PlusOne(s *Counter) {
	s.A++
	s.B++
}

func BenchmarkWithoutPool(b *testing.B) {
	var s *Counter
	s = &Counter{A: 1, B: 2}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			// b.StopTimer()
			PlusOne(s)
			// b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			p = pool.Get().(*Counter)
			p.A = 1
			p.B = 2
			// b.StopTimer()
			PlusOne(p)
			// b.StartTimer()
			pool.Put(p)
		}
	}
}
