package logic

import "testing"

func BenchmarkType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type(0, 0, 0)
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name(0, 0, 0)
	}
}

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Start(0, 0, 0)
	}
}
