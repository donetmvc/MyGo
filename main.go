package main

import (
	_ "eland-eoas-service/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"eland-eoas-service/models"
)

func init() {
	models.InitDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		orm.Debug = true
	}
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
