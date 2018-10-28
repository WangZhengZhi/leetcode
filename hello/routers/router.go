package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	//注意；当实现自定义的get方法，请求不会访问默认方法
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{},"get:ShowIndex;post:HandleLogin")
	beego.Router("/add", &controllers.MainController{},"get:ShowArtical;post:HandleArtical")
}
