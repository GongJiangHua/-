package controllers

import (
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"github.com/astaxie/beego"
)

type SendSmsController struct {
	beego.Controller
}

/**
	发送短信验证码
 */
func (s *SendSmsController)Post()  {
	var smsLogin models.SmsLogin
	err := s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("发送验证码解析数据失败，请重试！"
		return)
	}
	phone := smsLogin.Phone
	code := utils.GenRandCode(6)//返回一个6位的随机数
	utils.SendSms(phone,code)
}