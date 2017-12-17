package main

import (
	_ "eland-eoas-service/routers"
	"github.com/astaxie/beego"
	"eland-eoas-service/models"

	"fmt"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	error := models.InitDB("mysql", "")
	if error != nil {
		fmt.Println("Init db error.", error)
	}

	beego.Run()
}
