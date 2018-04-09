package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "主页-"+beego.AppConfig.String("title")
	c.Data["MyHome"] = true
	c.TplName = "home.html"

	c.Data["MyLogin"] = checkAccount(c.Ctx)
}

