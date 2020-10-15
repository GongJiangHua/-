package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName="login.html"
}
func (l *LoginController) Post()  {
	//1、解析客户端用户提交的登录数据
	var User models.User
	fmt.Println("正在解析。。。")
	err:=l.ParseForm(&User)
	if err!=nil {
		fmt.Println(err)
		l.Ctx.WriteString("登录页面数据解析错误，请重试！！")
		return
	}
	//2、根据解析到的数据，执行数据库查询操作
	//3、判断数据库查询结果
	u,err:=User.QueryUser()
	if err!=nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试")
		return
	}
	fmt.Println(u)
		//l.Ctx.WriteString("欢迎来到首界面！！")
	//4、根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone//动态数据设置
	l.TplName = "home.html"
}
