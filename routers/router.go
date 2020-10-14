package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router: 路由
    beego.Router("/", &controllers.MainController{})
	//用户注册接口
    beego.Router("/register",&controllers.RegisterControllers{})
	//用户登录的接口
    beego.Router("/login",&controllers.LoginControllers{})
	//请求直接登录的页面
    beego.Router("/login.html",&controllers.LoginControllers{})
	//用户上传文件的功能
    beego.Router("/home",&controllers.UploadFileControllers{})
}
