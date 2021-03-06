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
    //查看认证数据证书页面
    beego.Router("/cert_detail.html",&controllers.CertDetailController{})
    //用户实名认证
    beego.Router("/user_kyc",&controllers.UserKycController{})
    //跳转验证码登录页面
    beego.Router("/login_sms.html",&controllers.LoginSmsController{})
    //发送验证码
    beego.Router("/send_sms",&controllers.SendSmsController{})
    beego.Router("/login_sms",&controllers.SendSmsController{})
}
