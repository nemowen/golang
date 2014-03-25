package controllers

import (
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	//set templater
	this.Data["PageTitle"] = "Topic"
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	beego.Debug("debug topic")
	this.TplNames = "topic.html"

}
