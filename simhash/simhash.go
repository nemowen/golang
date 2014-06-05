package main

import (
	//"encoding/binary"
	"fmt"
	"github.com/mfonda/simhash"
	"strconv"
)

func main() {

	var docs = [][]byte{
		[]byte("Z2X6629402"),
		[]byte("q2X6629302"),
		[]byte("ZyX6629302"),
	}

	uint64s, _ := strconv.ParseUint("1001100110111100011011010111101101100111100111000110111010010", 2, 64)
	fmt.Printf("64:%x\n", uint64s)

	hashes := make([]uint64, len(docs))
	for i, d := range docs {
		hashes[i] = simhash.Simhash(simhash.NewWordFeatureSet(d))
		fmt.Printf("Simhash of %s: %x b:%b\n", d, hashes[i], hashes[i])
	}

	q, e := strconv.ParseUint("13378daf6cf38dd2", 16, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	q2, e := strconv.ParseUint("f597e9511bbc518d", 16, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Printf("simhash:%d\n", simhash.Compare(q, q2))

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
}
