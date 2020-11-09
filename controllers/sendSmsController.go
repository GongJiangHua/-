package controllers

import (
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
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
		s.Ctx.WriteString("发送验证码解析数据失败，请重试！")
		return
	}
	phone := smsLogin.Phone
	code := utils.GenRandCode(6)//返回一个6位的随机数
	result,err := utils.SendSms(phone,code,utils.SMS_TLP_REGISTER)
	if err != nil{
		s.Ctx.WriteString("发送验证码失败，请重试！")
		return
	}
	fmt.Println("发送成功")
	if len(result.Bizid)==0 {
		s.Ctx.WriteString(result.Message)
		return
	}
	smsRecord := models.SmsRecord{
		BizId:     result.Bizid,
		Phone:     phone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStamp: time.Now().Unix(),
	}
	_,err = smsRecord.SaveSmsRecord()
	if err != nil {
		s.Ctx.WriteString("发送短信的信息保存失败，请重试！")
		return
	}
	//保存成功
	s.Data["Phone"] =smsLogin.Phone
	s.Data["BizId"]=smsRecord.BizId
	s.TplName="login.html"
}