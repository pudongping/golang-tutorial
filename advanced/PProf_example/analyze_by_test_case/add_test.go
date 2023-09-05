package analyze_by_test_case

import (
	"testing"
)

func TestAdd(t *testing.T) {
	_ = Add("go-programming-tour-book")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("go-programming-tour-book")
	}
}
