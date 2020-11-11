package models

import (
	"HelloBeego01/db_mysql"
	"fmt"
)

type SmsRecord struct {
	BizId string
	Phone string
	Code string
	Status string
	Message string
	TimeStamp int64
}

func (s SmsRecord)SaveSmsRecord() (int64,error) {
	res,err := db_mysql.DB.Exec("insert into sms_record(biz_id,phone,code,status,message,timestemp) values (?,?,?,?,?,?)",
		s.BizId,s.Phone,s.Code,s.Status,s.Message,s.TimeStamp)
	fmt.Println(s.BizId,s.Phone,s.Code,s.Status,s.Message,s.TimeStamp)
	if err != nil {
		return -1,err
	}
	return res.RowsAffected()
}

/**
	根据用户提交的手机号和短信验证码查询验证码是否正确
 */
