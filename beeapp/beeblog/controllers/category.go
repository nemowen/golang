package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	//set templater
	this.Data["PageTitle"] = "Category"
	this.Data["IsCategory"] = true
	this.TplNames = "category.html"

}
