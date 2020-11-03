package controllers

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

/**
	钙get方法用于处理浏览器的get请求，往查看证书详情页面跳转
 */
func (c *CertDetailController) Get() {
	//解析和接收前端页面传递的数据cert_id
	cert_id := c.GetString("cert_id")
	//2、到区块链上查询区块数据
	block,err := blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if err!= nil {
		c.Ctx.WriteString("抱歉，查询链上数据失败，请重试！！")
		return
	}
	if block == nil{//遍历整条区块链，但是未找到数据
		c.Ctx.WriteString("抱歉，未找到链上数据！！")
		return
	}
	fmt.Println("查询到区块的高度：",block.Height)
	//反序列化
	certRecord,err := models.DeserializeCertRecord(block.Data)
	certRecord.CertIdHex= strings.ToUpper(string(certRecord.CertId))
	certRecord.CertHashHex= string(certRecord.CertHash)
	certRecord.CertTimeFormat=utils.TimeFormat(certRecord.CertTime,utils.TIME_FORMAT_ONE)
	//结构体
	c.Data["CertRecord"]=certRecord
	//跳转证书详情页面
	c.TplName="cert_detail.html"
}