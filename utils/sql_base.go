package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


//链接数据库
//dbhost数据库地址
//dbport数据库端口
//dbuser数据库用户名
//dbpassword数据库密码
//dbname数据库名称
func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	
	orm.RegisterDataBase("default", "mysql", dsn)

	//注册生成相关的表
	orm.RegisterModel()
}
