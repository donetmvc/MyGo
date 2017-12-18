package controllers

import (
	"github.com/astaxie/beego"
	"eland-eoas-service/models"

	"encoding/json"
)
// Login API
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description Login with userName and password
// @Success 200 {object} models.LoginInfo
// @Param   userName     path   string true       "user name"
// @Param   password    path   int true       "user password"
// @Failure 400 no enough input
// @Failure 500 login error
// @router /:userName/:password [get]
func (l *LoginController) GetLogin() {
	username := l.Ctx.Input.Param(":userName")
	password := l.Ctx.Input.Param(":password")
	if username != "" && password != "" {
		info, err := models.Login(username, password)

		apiResult := new(models.APIResult)
		apiError := new(models.APIError)

		if err != nil {
			apiError.Message = err.Error()
		} else {
			infoJson, jsonError := json.Marshal(info)
			if jsonError != nil {
				apiError.Message = "json convert error"
				
			} else {
				apiResult.Success = true
				apiResult.Result = string(infoJson)
			}
		}
		apiResult.Error = apiError
		// if apiError.Message != "" {
		// 	msgBody, _:= json.Marshal(apiResult)
		// 	l.CustomAbort(500, string(msgBody))
		// }
		l.Data["json"] = apiResult
	}
	l.ServeJSON()
}