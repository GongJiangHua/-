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

func (l *LoginControllers) Post()  {
	//1、解析客户端用户提交的登录数据
	var loginUser models.User
	fmt.Println("正在解析。。。")
	err:=l.ParseForm(&loginUser)
	if err!=nil {
		fmt.Println(err)
		l.Ctx.WriteString("登录页面数据解析错误，请重试！！")
		return
	}
	//2、根据解析到的数据，执行数据库查询操作
	//3、判断数据库查询结果
	user,err :=db_mysql.QueryUser(loginUser)
	if err!=nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("用户不存在，请注册！！")
		return
	}

	//4、根据查询结果返回客户端相应的信息或者页面跳转



}
