package models

import (
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

	"time"
	"fmt"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", beego.AppConfig.String("mysqlConn"))
	if err != nil {
		fmt.Print(err.Error())
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.ShowSQL(true)
}

// func InitDB() {
// 	var err error
// 	engine, err = xorm.NewEngine("mysql", beego.AppConfig.String("mysqlConn"))
// 	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
// 	engine.Sync2(new(LoginInfo), new(Detail))
// }