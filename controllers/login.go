package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type LoginController struct {
	BaseController
}


//用来进行添加用户的操作
func (u *LoginController) AddUser() {
	u.SetSession("liukai","234234234234324")
	sess := u.GetSession("liukai")
	beego.Info(sess)
	u.Data["json"] = map[string]string{"uid": "123123","name":"你好"}
	u.ServeJSON()
}

func (u *LoginController) GetUser() {
	sess := u.GetSession("liukai")
	if
}

