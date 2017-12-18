package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)


func InitDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlConn"))

	orm.RegisterModel(new(LoginInfo), new(Detail))
}