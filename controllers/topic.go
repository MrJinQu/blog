package controllers

import (
	"github.com/astaxie/beego"
	"test/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get()  {
	c.Data["MyTopic"] = true
	c.Data["Title"] = "文章-"+beego.AppConfig.String("title")
	c.Data["MyLogin"] = checkAccount(c.Ctx)
	c.TplName="topic.html"
}
func (c *TopicController) Post()  {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login",302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")

	var err error
	err = models.AddTopic(title,content)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic",302)
}

func (c *TopicController) Add()  {
	c.TplName="topic-add.html"
	c.Data["Title"] = "添加文章"
}