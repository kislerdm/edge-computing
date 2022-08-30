package logic

import "testing"

func BenchmarkTypeBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type(0, 0, 0)
	}
}

func BenchmarkNameBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name(0, 0, 0)
	}
}

func BenchmarkStartBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Start(0, 0, 0)
	}
}

func BenchmarkTypeWhite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type(255, 255, 255)
	}
}

func BenchmarkNameWhite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name(255, 255, 255)
	}
}

func BenchmarkStartWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Start(255, 255, 255)
	}
}
