package controllers

import (
	"github.com/astaxie/beego"
)

type BaopengController struct {
	beego.Controller
}
func (c *BaopengController) Get() {
	c.TplName = "baopeng.html"
}
