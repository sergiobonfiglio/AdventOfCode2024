package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2_v1(input)
	}
}

func BenchmarkPart2_BronKerbosch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2_BronKerbosch(input)
	}
}
