package classtest

import (
	"testing"
)

func BenchmarkDeal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deal()
	}
}
