package main

import (
	"fmt"
	"github.com/mfonda/simhash"
	"strconv"
)

func main() {

	var docs = [][]byte{
		[]byte("34 23 56 567 878"),
		[]byte("33 21 53 560 871"),
		[]byte("34 22 56 567 878"),
	}

	hashes := make([]uint64, len(docs))
	for i, d := range docs {
		hashes[i] = simhash.Simhash(simhash.NewWordFeatureSet(d))
		fmt.Printf("Simhash of %s: %x\n", d, hashes[i])
	}
	q, e := strconv.ParseUint("8329707b4eb870c", 16, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	q2, e := strconv.ParseUint("8329707b4eb8707", 16, 64)

	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Printf("qwqw:%d\n", simhash.Compare(q, q2))

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
}
