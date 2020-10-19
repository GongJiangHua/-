package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router: 路由
    beego.Router("/", &controllers.MainController{})
	//用户注册接口
    beego.Router("/register",&controllers.RegisterController{})
	//用户登录的接口
    beego.Router("/login",&controllers.LoginController{})
	//请求直接登录的页面
    beego.Router("/login.html",&controllers.LoginController{})
	//用户上传文件的功能
    beego.Router("/home",&controllers.UploadFileController{})
    //用户上传后跳转到信息界面
    beego.Router("/record",&controllers.UploadFileController{})
    //用户新增存证记录跳转home页面
    beego.Router("/home.html",&controllers.UploadFileController{})
}
