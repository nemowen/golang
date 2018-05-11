package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	message.Set(language.English, "我有 %d 个苹果",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))
}

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("我有 %d 个苹果", 1)
	p.Println()
	p.Printf("我有 %d 个苹果", 5)
	p.Println()
	p.Printf("我有 %d 个苹果", 2)
	p.Println()
}
