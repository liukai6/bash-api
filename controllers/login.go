package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type LoginController struct {
	beego.Controller
}


func (u *LoginController) Get() {
	beego.Info("部署")
	u.Data["json"] = map[string]string{"uid": "123123","name":"你好"}
	u.ServeJSON()
}


