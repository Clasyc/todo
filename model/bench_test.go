package model

import (
	"strings"
	"testing"
)

func BenchmarkString1(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str += "benchmark"
	}
}

func BenchmarkString2(b *testing.B) {
	str := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str[i] = "benchmark"
	}
	_ = strings.Join(str, "")
}
