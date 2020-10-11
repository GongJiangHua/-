package controllers

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterControllers struct {
	beego.Controller
}

func (r *RegisterControllers) Post() {
	//1.解析用户提交的请求数据
	var User models.User
	err :=r.ParseForm(&User)
	if err !=nil {
		r.Ctx.WriteString("抱歉，数据解析失败，请重试！！" )
		return
	}
	//2.将解析到的数据保存到数据库中
	row,err :=db_mysql.AddUser(User)
	if err !=nil{
		r.Ctx.WriteString("数据库存储用户信息失败！！")
		return
	}
	fmt.Println(row)
	//3.将处理结果返回给客户端浏览器
	//3.1如果成功，跳转到登录页面 template模板
	if row!=-1 {
		r.TplName="login.html"
		return
	}else {
	//3.2如果错误，跳转到错误页面

	}



}
