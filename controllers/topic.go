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
	topics,err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err.Error())
	}else {
		c.Data["Topic"] = topics
	}
}
func (c *TopicController) Post()  {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login",302)
		return
	}
	//tid := c.Input().Get("tid")
	//title := c.Input().Get("title")
	//content := c.Input().Get("content")
	//
	//var err error
	//if len(tid)==0{
	//err = models.AddTopic(title,content)
	//}else{
	//	err =models.ModifyTopic(tid,title,content)
	//}
	//if err != nil {
	//	beego.Error(err)
	//}
	//c.Redirect("/topic",302)
}

func (c *TopicController) Add()  {
	c.TplName="topic-add.html"
	c.Data["Title"] = "添加文章"
	c.Data["MyLogin"] = checkAccount(c.Ctx)
}

func (c *TopicController) Post()  {
	c.Data["MyLogin"] = checkAccount(c.Ctx)
	c.Data["Title"] = beego.AppConfig.String("title")
	c.TplName="topic-view.html"
	if !checkAccount(c.Ctx){
		c.Redirect("/login",302)
		return
	}
	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	content := c.Input().Get("content")

	var err error
	if len(tid)==0 {
		err = models.AddTopic(title,content)
	}else {
		err = models.Modifytopic(tid,title,content)
	}

	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic",302)
}


func (c *TopicController) View() {
	c.Data["MyLogin"] = checkAccount(c.Ctx)
	c.Data["Title"] = beego.AppConfig.String("title")
	c.TplName="topic-view.html"

	topic , err := models.GetTopic(c.Ctx.Input.Params["0"])

	i {

	}
}
func (c *TopicController) Modify () {
	c.Data["MyLogin"] = checkAccount(c.Ctx)
	c.TplName="topic-modify.html"
	tid := c.Input().Get("tid")
	topic , err := models.GetTopic(tid)
	if err!= nil {
		beego.Error(err)
		c.Redirect("/",302)
		return
	}
	c.Data["Topic"]=topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
}