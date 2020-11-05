package utils

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
)

/**
	该函数用于发送一条短信
	phone ：电话，用于接收验证码的手机
	code:发送验证码的数字
	templateType：模板类型
 */
type SmsCode struct {
	Code string `json:"code"`
}
type SmsResult struct {
	Bizid string
	Code string
	Message string
	Requestid string
} 

func SendSms(phone string,code string,templateType string)(*SmsResult,error) {
	config := beego.AppConfig
	AccessKey := config.String("accessKey")
	AcceaaKeySecret:=config.String("acceaaKeySecret")
	client,err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",AccessKey,AcceaaKeySecret)
	if err != nil {
		return nil,err
	}
	//batch 批量，批次
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone//指定要发送的那个目标手机号
	request.SignName = "线上餐厅"//签名信息
	request.TemplateCode = templateType //指定短信模板
	//{"code":"XXXXXX"}Json格式
	smsCode := SmsCode{Code:code}
	smsbytes,_ := json.Marshal(smsCode)
	request.TemplateParam = string(smsbytes)//指定要发送的验证码
	
	response,err := client.SendSms(request)
	if err != nil {
		return nil,err
	}
	//Biz : business 商业，业务
	smsResult := &SmsResult{
		Bizid:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		Requestid: response.RequestId,
	}
	return smsResult,nil
}

//func GenRandCode(width int)string  {
//
//}