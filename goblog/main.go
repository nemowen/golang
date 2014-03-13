package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this MainController) Get() {
	this.Ctx.WriteString("Hello Nemo! Good Night!")
}

func main() {
	beego.Router("/index", &MainController{})
	beego.Run()
}
