package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

/*
	手动加载消息
*/
type entry struct {
	tag, key string
	msg      interface{}
}

var entries = [...]entry{
	{"en", "Hello World", "Hello World"},
	{"zh", "Hello World", "你好世界"},
	{"en", "%d task(s) remaining!", plural.Selectf(1, "%d",
		"=1", "One task remaining!",
		"=2", "Two tasks remaining!",
		"other", "[1]d tasks remaining!",
	)},
	{"zh", "%d task(s) remaining!", plural.Selectf(1, "%d",
		"=1", "剩余一项任务！",
		"=2", "剩余两项任务！",
		"other", "剩余 [1]d 项任务！",
	)},
}

func init() {
	for _, e := range entries {
		tag := language.MustParse(e.tag)
		switch msg := e.msg.(type) {
		case string:
			message.SetString(tag, e.key, msg)
		case catalog.Message:
			message.Set(tag, e.key, msg)
		case []catalog.Message:
			message.Set(tag, e.key, msg...)
		}
	}
}

func main() {
	p := message.NewPrinter(language.Chinese)

	p.Printf("Hello World")
	p.Println()
	p.Printf("%d task(s) remaining!", 2)
	p.Println()

	p = message.NewPrinter(language.English)
	p.Printf("Hello World")
	p.Println()
	p.Printf("%d task(s) remaining!", 2)

}
