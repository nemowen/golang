package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gotest/beeapp/beeblog/controllers"
	"gotest/beeapp/beeblog/models"
	_ "gotest/beeapp/beeblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}
