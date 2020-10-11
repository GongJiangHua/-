package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//路由
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.RegisterControllers{})
}
