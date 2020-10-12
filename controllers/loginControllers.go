package controllers

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}

func (l *LoginControllers) Get()  {
	var loginUser models.Userlogin
	fmt.Println("正在解析。。。")
	err:=l.ParseForm(&loginUser)
	if err!=nil {
		l.Ctx.WriteString("登录页面数据解析错误，请重试！！")
		return
	}
	pwd,err :=db_mysql.QueryUser(loginUser)
	fmt.Println(pwd)
}
