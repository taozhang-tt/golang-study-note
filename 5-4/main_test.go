package main

import "testing"

func BenchmarkMain2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main2()
	}
}

func BenchmarkMain3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main3()
	}
}
