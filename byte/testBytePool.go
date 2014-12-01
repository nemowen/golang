package gotest

func BenchmarkBytepoolAllocation(b *testing.B) {
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
