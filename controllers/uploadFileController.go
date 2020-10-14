package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

/*
该控制器结构体用于处理文件上传的功能
*/
type UploadFileControllers struct {
	beego.Controller
}

func (u *UploadFileControllers) Post() {
	//用户上传的自定义标题
	title :=u.Ctx.Request.PostFormValue("upload_title")
	//用户上传的文件
	file,header,err :=u.GetFile("file")
	if err != nil {
		u.Ctx.WriteString("抱歉，文件保存失败，请重试！！")
		return
	}
	fmt.Println("自己定义的文件标题：",title)
	fmt.Println("文件名称：",header.Filename)
	fmt.Println("文件大小：",header.Size)
	fmt.Println("上传的文件：",file)
	//u.Ctx.WriteString("文件上传成功！！")
	//eg：支持jpg,png类型，不支持jpeg，gif类型
	//文件名： 文件名 + "." + 扩展名
	isJpg := strings.HasSuffix(header.Filename,".jpg")
	isPng := strings.HasSuffix(header.Filename,".png")
	if !isJpg && !isPng {
		u.Ctx.WriteString("上传的文件类型不符合规范，请检查重新上传！！")
		return
	}
	//文件的大小 200kb
	config := beego.AppConfig
	filesize,err := config.Int64("file_size")
	if header.Size/1024 > filesize {
		u.Ctx.WriteString("文件太大，请重新上传符合大小的文件！！")
		return
	}
	u.Ctx.WriteString("文件上传成功！！")
}