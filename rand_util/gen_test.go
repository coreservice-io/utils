package rand_util

import (
	"testing"
)

func Benchmark_randStr(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenRandStr(i)
	}
}
