package routers

import (
	"bash-api/controllers"
	"github.com/astaxie/beego"
)

func apiRouter()  {
	beego.Router("/api/adduser",&controllers.LoginController{},"post:AddUser")

}

