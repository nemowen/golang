package temptest

import (
	l "container/list"
	"testing"
)

var (
	name = "viney"
)

func BenchmarkList(b *testing.B) {
	names := l.New()
	b.N = 1000000000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = names.PushFront(name)
	}
	b.StopTimer()
}

func BenchmarkSlice(b *testing.B) {
	names := []string{}
	for i := 1; i < b.N; i++ {
		names = append(names, name)
	}
}
