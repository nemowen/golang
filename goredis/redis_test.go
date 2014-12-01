package goredis

import (
	"strconv"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		redisClient.Set("key"+strconv.Itoa(i), []byte("valeu"+strconv.Itoa(i)))
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		redisClient.Get("key" + strconv.Itoa(i))
	}
}
