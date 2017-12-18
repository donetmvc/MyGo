package controllers

import (
	"github.com/astaxie/beego"
	"eland-eoas-service/models"
)
// Login API
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description Login with username and password
// @Success 200 {object} models.LoginInfo
// @Param   username     query   string false       "user name"
// @Param   password    query   int false       "user password"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router api/v1/login [get]
func GetLogin(l *LoginController) {
	username := l.Ctx.Input.Param(":userName")
	password := l.Ctx.Input.Param(":password")
	if username != "" && password != "" {
		info, err := models.ExtendLogin(username, password)
		if err != nil {
			l.Data["json"] = err.Error()
		} else {
			l.Data["json"] = info
		}
	}
	l.ServeJSON()
}