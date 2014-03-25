package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["PageTitle"] = "Login"
	this.TplNames = "login.html"
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("upwd", "", -1, "/")
	}
	beego.Debug("okokokokokoko")
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	upwd := this.Input().Get("upwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("upwd") == upwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("upwd", upwd, maxAge, "/")
	} else {
		this.Redirect("/login?err=20112", 301)
		return
	}

	this.Redirect("/", 301)
	return
}

func checkAccount(this *context.Context) bool {
	v, err := this.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := v.Value

	v, err = this.Request.Cookie("upwd")
	if err != nil {
		return false
	}
	upwd := v.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("upwd") == upwd
}
