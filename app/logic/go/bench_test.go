package main

import (
	"testing"

	"edgecomputing/logic"
)

func BenchmarkTypeBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logic.Type(0, 0, 0)
	}
}

func BenchmarkNameBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logic.Name(0, 0, 0)
	}
}

func BenchmarkStartBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t(0, 0, 0)
		n(0, 0, 0)
		getNAddress()
		getNLen()
	}
}

func BenchmarkTypeWhite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logic.Type(255, 255, 255)
	}
}

func BenchmarkNameWhite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logic.Name(255, 255, 255)
	}
}

func BenchmarkStartWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t(255, 255, 255)
		n(255, 255, 255)
		getNAddress()
		getNLen()
	}
}
