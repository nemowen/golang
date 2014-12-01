package main

import "fmt"

type USB interface {
	Name()
	Connecter
}

type Connecter interface {
	Connect()
}

type Keybord struct {
	name string
}

func (k Keybord) Name() {
	fmt.Println(k.name)
}

func (k Keybord) Connect() {
	fmt.Println("connected to ", k.name)
}

func main() {
	keybord := Keybord{"双飞燕键盘"}
	keybord.Name()
	keybord.Connect()
}
