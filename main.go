package main

import (
	_"test/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/models"
)

func init()  {
	models.RegisterDB()
}
func main()  {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	orm.RunCommand()
	beego.Run()
}