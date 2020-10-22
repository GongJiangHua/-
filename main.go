package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0 := blockchain.CreateGenesisBlock()
	fmt.Println(block0)
	fmt.Printf("block0的hash值:%x\n",block0.Hash)
	block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,nil)
	fmt.Printf("block1的hash值:%x\n",block1.Hash)
	fmt.Printf("block0的prev hash值:%x\n",block1.PrevHash)
	blockchain.UnSerialize()
	//1、序列化:
		//将数据从内存中形式转化为可以持久化存储在硬盘上或者在网络上传播
	//block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("abc"))
	//fmt.Println("block1的hash:",block1.Hash)
	//blockchain.NewPow(block1)
	//连接数据库
	db_mysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

