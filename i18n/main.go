package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	message.SetString(language.Chinese, "%s went to %s.", "%s去了%s。")
	message.SetString(language.AmericanEnglish, "%s went to %s.", "%s is in %s.")
	message.SetString(language.Chinese, "%s has been stolen.", "%s被偷走了。")
	message.SetString(language.AmericanEnglish, "%s has been stolen.", "%s has been stolen.")
	message.SetString(language.Chinese, "How are you?", "你好吗?.")
}

func main() {
	p := message.NewPrinter(language.Chinese)
	p.Printf("%s went to %s.", "彼得", "英格兰")
	fmt.Println()
	p.Printf("%s has been stolen.", "宝石")
	fmt.Println()

	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("%s went to %s.", "Peter", "England")
	fmt.Println()
	p.Printf("%s has been stolen.", "The Gem")
}
