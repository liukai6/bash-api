package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)

//进行过滤
//此处需要注意的内容是引入文件的路径是否正确
func init() {
	var FinishRouter = func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Add("MinDoc-Version", "v1.00")
		ctx.ResponseWriter.Header().Add("MinDoc-Site", "https://www.iminho.me")
		ctx.ResponseWriter.Header().Add("X-XSS-Protection", "1; mode=block")
	}

	//var StartRouter = func(ctx *context.Context) {
	//	sessionId := ctx.Input.Cookie(beego.AppConfig.String("sessionname"))
	//	if sessionId != "" {
	//		//sessionId必须是数字字母组成，且最小32个字符，最大1024字符
	//		if ok, err := regexp.MatchString(`^[a-zA-Z0-9]{32,512}$`, sessionId); !ok || err != nil {
	//			panic("401")
	//		}
	//	}
	//}
	//beego.InsertFilter("/*", beego.BeforeStatic, StartRouter, false)
	beego.InsertFilter("/*", beego.BeforeRouter, FinishRouter, false)
}
