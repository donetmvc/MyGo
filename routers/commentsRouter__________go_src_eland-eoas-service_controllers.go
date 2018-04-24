package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["eland-eoas-service/controllers:LoginController"] = append(beego.GlobalControllerRouter["eland-eoas-service/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["eland-eoas-service/controllers:LoginController"] = append(beego.GlobalControllerRouter["eland-eoas-service/controllers:LoginController"],
		beego.ControllerComments{
			Method: "GetLogin",
			Router: `/:userName/:password`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
