package gotest

import (
	"fmt"

	"github.com/viki-org/bytepool"
	"testing"
)

func BenchmarkBytepoolAllocation(b *testing.B) {
	var pool = bytepool.New(8196, 32768)
	for j := 0; j < b.N; j++ {
		buffer := pool.Checkout()
		buffer.BeginArray()
		for i := 0; i < 20000; i++ {
			buffer.WriteInt(j)
		}
		buffer.EndArray()
		fmt.Sprintf(buffer.String())
	}
}
