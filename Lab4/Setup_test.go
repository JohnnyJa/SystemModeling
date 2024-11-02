package Lab4

import (
	"testing"
)

func Benchmark1(b *testing.B) {
	m := Setup3(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Simulate()
	}
}
