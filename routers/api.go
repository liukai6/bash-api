package routers

import (
	"base-api/controllers"
	"github.com/astaxie/beego"
)

func apiRouter()  {
	beego.Router("/api",&controllers.LoginController{})
}

