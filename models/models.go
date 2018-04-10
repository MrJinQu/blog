package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"os"
	"path"
	"time"
	"strconv"
	"github.com/astaxie/beego"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id               int
	Uid              int
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int	`orm:"index"`
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int
	ReplayLastUserId int
}


func RegisterDB() {

	dbuser := beego.AppConfig.String("dbuser")
	dbpass := beego.AppConfig.String("dbpass")
	dbip := beego.AppConfig.String("dbip")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")

	_DB_NAME := dbuser + ":"+ dbpass + "@tcp("+ dbip + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"

	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}



	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", _DB_NAME, 5,30)
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title",name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func DelCategory(id string) error  {
	cid, err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o :=orm.NewOrm()
	cate := &Category{Id:cid}
	_,err=o.Delete(cate)
	return err
}
func GetAllCategories() ([]*Category, error)  {
	o := orm.NewOrm()

	cates := make([]*Category,0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates,err
}
func AddTopic(title,content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
	}
	_,err := o.Insert(topic)
	return err
}