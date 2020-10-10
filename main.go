package main

import (
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	//连接数据库
	db_mysql.Connect()
	beego.Run()
}

