package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func init() {
	message.Set(
		language.English, "你迟到了 %d 分钟!",
		catalog.Var("m", plural.Selectf(1, "%d",
			"one", "minute",
			"other", "minutes")),
		catalog.String("You are %[1]d ${m} late!"),
	)
}

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("你迟到了 %d 分钟!", 1)
	p.Println()
	p.Printf("你迟到了 %d 分钟!", 10)
	p.Println()
}
