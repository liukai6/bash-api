package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type LoginController struct {
	beego.Controller
}


//用来进行添加用户的操作
func (u *LoginController) AddUser() {

	beego.Info("部署")
	u.Data["json"] = map[string]string{"uid": "123123","name":"你好"}
	u.ServeJSON()
}


