package controllers

import (
	"DataCertPlatform/models"
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
	var smsLogin models.SmsLogin
	err := l.ParseForm(&smsLogin)
	if err != nil {
		l.Ctx.WriteString("抱歉，验证码登陆失败，请重试")
		return
	}
	//将用户提交的信息和数据库里的进行比较

}