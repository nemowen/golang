package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var (
	list PList
)

type Product struct {
	title    string
	price    float64
	quantity int64
}

func (p *Product) String() string {
	return p.title + " --> " +
		strconv.FormatFloat(p.price, 'f', 1, 64) + " --> " +
		strconv.FormatInt(p.quantity, 10)
}

type PList []*Product

func (pl *PList) Add(p *Product) {
	*pl = append(*pl, p)
}

func NewProduct(s []string) *Product {
	p, _ := strconv.ParseFloat(s[1], 32)
	q, _ := strconv.ParseInt(s[2], 10, 64)
	return &Product{
		title:    s[0],
		price:    p,
		quantity: q,
	}
}

// parsing csv-files

// the file contents: products.txt
//
// The Abc of Go;25.5;1500
// Functional Programming with Go;56;280
// Go fro It;45.9;356
// The Go way;55;500
func main() {
	file, err := os.Open("C:/Users/nemowen/products.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	allf, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, v := range allf {
		list.Add(NewProduct(v))
	}

	fmt.Println(list)

}
