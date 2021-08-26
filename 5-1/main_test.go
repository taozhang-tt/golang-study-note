package main

import (
	"testing"
)

// go test -bench=. -benchmem
// BenchmarkTest1-4   	   5704	   199652 ns/op	 530339 B/op	    999 allocs/op
// BenchmarkTest2-4   	  99636	    12723 ns/op	   1024 B/op	      1 allocs/op
// BenchmarkTest3-4   	 117040	    10187 ns/op	   2048 B/op	      2 allocs/op
// BenchmarkTest4-4   	  91966	    11239 ns/op	   3248 B/op	      6 allocs/op

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test1()
	}
}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}

func BenchmarkTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test3()
	}
}

func BenchmarkTest4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test4()
	}
}
