package controllers

import (
	"github.com/astaxie/beego"
)

type LoginSmsController struct {
	beego.Controller
}

func (l *LoginSmsController)Get()  {
	l.TplName="login_sms.html"
}
/**
	发送短信验证码功能
 */
func (l *LoginSmsController)Post()  {

}