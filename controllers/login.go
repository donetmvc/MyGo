package controllers

import (
	"github.com/astaxie/beego"
	"eland-eoas-service/models"

	"encoding/json"
)

var (
	apiResult *models.APIResult
	apiError *models.APIError
)

func init() {
	apiResult = new(models.APIResult)
	apiError = new(models.APIError)
}

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
		if err != nil {
			apiError.Message = err.Error()
		} else {
			infoJson, jsonError := json.Marshal(&info)
			if jsonError != nil {
				apiError.Message = "json convert error"
				
			} else {
				apiResult.Success = true
				apiResult.Result = string(infoJson)
			}
		}
		apiResult.Error = apiError
		l.Data["json"] = apiResult
	}
	l.ServeJSON()
}


// @Title CreateUser
// @Description create users
// @Param	body		body 	models.LoginInfo	true		"body for user content"
// @Success 200 {int} models.LoginInfo.Id
// @Failure 403 body is empty
// @router / [post]
func (l *LoginController) Register() {
	var userInfo models.LoginInfo
	json.Unmarshal(l.Ctx.Input.RequestBody, &userInfo)
	id, err := models.CreateUser(&userInfo)

	if err != nil {
		apiError.Message = err.Error()
	} else {
		apiResult.Success = true
		apiResult.Result = `{"id":` + string(int(id)) + `}`
	}

	l.Data["json"] = apiResult
	l.ServeJSON()
}