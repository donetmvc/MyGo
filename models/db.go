package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego/config"

	"time"
	"fmt"
	"errors"
)

var engine *xorm.Engine

func InitDB(sqlDriver, conn string) error {
	var err error
	appConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		return err
	}
	//

	sqlconn := appConfig.String("mysqlConn")
	fmt.Println(sqlconn)

	engine, err = xorm.NewEngine(sqlDriver, sqlconn)
	if err != nil {
		return err
	}

	engine.ShowSQL(true)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	return engine.Sync2(new(UserInfo))
}

func Query(conn string) (u *UserInfo, err error) {
	result, err := engine.Query(conn)

	if(err != nil) {
		return nil, err
	}

	return new(UserInfo), nil
}