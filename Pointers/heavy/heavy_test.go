package heavy

import "testing"

var sink int

func BenchmarkByValue(b *testing.B) {
	h := Heavy{}
	for i := 0; i < b.N; i++ {
		sink = ByValue(h)
	}
}

func BenchmarkByPointer(b *testing.B) {
	h := Heavy{}
	for i := 0; i < b.N; i++ {
		sink = ByReferenc(&h)
	}
}
