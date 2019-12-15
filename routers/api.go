package routers

import (
	"bash-api/controllers"
	"github.com/astaxie/beego"
)

func apiRouter()  {
	beego.Router("/api/adduser",&controllers.LoginController{},"get:AddUser")
	beego.Router("/api/getuser",&controllers.LoginController{},"get:GetUser")

}

