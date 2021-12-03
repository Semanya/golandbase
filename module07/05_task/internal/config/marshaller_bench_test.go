package config

import (
	"module07/internal/convertor"
	"testing"
)

var simple = Config{
	Name:        "aaa",
	Host:        "bbb",
	Port:        123,
	Environment: map[string]string{"bacon": "delicious"},
	Volumes:     []string{"a", "b", "c", "d", "e", "f", "g"},
}

func Benchmark2StructToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simple.StructToMap()
	}
}

func BenchmarkStructToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertor.StructToMap(simple)
	}
}
