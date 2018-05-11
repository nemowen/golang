package main

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("%d", currency.Symbol(currency.USD.Amount(0.1)))
	p.Println()
	p.Printf("%d", currency.NarrowSymbol(currency.CNY.Amount(1.6)))
	p.Println()
	p.Printf("%d", currency.ISO.Kind(currency.Cash)(currency.EUR.Amount(123.456)))
	p.Println()
}
