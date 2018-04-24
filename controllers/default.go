package controllers

import (
	"github.com/astaxie/beego"
	"test/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "主页-"+beego.AppConfig.String("title")
	c.Data["MyHome"] = true
	c.TplName = "home.html"
	c.Data["MyLogin"] = checkAccount(c.Ctx)
	topics,err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err.Error())
	}else {
		c.Data["Topic"] = topics
	}
}

