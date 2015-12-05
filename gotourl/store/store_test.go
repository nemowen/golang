package store

import (
	"fmt"
	"testing"
)

var store = NewURLStore()

func TestPutURL(t *testing.T) {
	if "0a" != store.Put("http://www.baidu.com") {
		t.Fatal("want 0a")
	}
}

func TestGetURL(t *testing.T) {
	if "http://www.baidu.com" != store.get("0a") {
		t.Fatal("want http://www.baidu.com")
	}

}

func BenchmarkPutURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		store.Put(fmt.Sprintf("http://www.baidu%v.com", i))
	}
	b.Logf("store.count() is :%v", store.Count())
}
