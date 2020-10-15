package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
该控制器结构体用于处理文件上传的功能
*/
type UploadFileController struct {
	beego.Controller
}

func (u *UploadFileController) Post1() {
	//用户上传的自定义标题
	title :=u.Ctx.Request.PostFormValue("upload_title")
	//用户上传的文件
	file,header,err :=u.GetFile("file")

	if err != nil {
		u.Ctx.WriteString("抱歉，文件保存失败，请重试！！")
		return
	}
	defer file.Close()
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

	//fromFile: 文件，
	//toFile: 要保存的文件路径，
	//权限的组成
	     //a : 文件所有者对文件的操作权限，读4、写2、执行1
	     //b ：文件所有者所在组的用户的操作权限，读4、写2、执行1
	     //c : 其他用户的操作权限，读4、写2、执行1

	     //eg: m文件，权限是：651
	     //判读题：文件所有者对该m文件有写权限（对）
	saveDir := "static/upload"
	//打开该文件
	f,err := os.Open(saveDir)
	if err!=nil {
		//创建文件夹
		err = os.Mkdir(saveDir,777)
		if err!=nil {
			u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试！！")
			return
		}
	}
	fmt.Println(f)

	//文件名：   文件路径 + 文件名 + "." + 文件拓展名
	saveName := saveDir + "/" + header.Filename
	fmt.Println("要保存的文件名：",saveName)
	err = u.SaveToFile("file",saveName)
	if err!=nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("文件保存失败，请重试！！")
	}
	u.Ctx.WriteString("文件上传成功！！")
}

/*
该post方法用于处理用户在客户端提交的文件
 */
func (u *UploadFileController)Post()  {
	//用户上传的自定义标题
	title :=u.Ctx.Request.PostFormValue("upload_title")
	//用户上传的文件
	file,header,err :=u.GetFile("file")
	fmt.Println(title)
	if err != nil {
		u.Ctx.WriteString("抱歉，文件保存失败，请重试！！")
		return
	}
	defer file.Close()
	//使用io包提供的方法保存文件
	saveFilePath  := "static/upload/"+header.Filename
	saveFile ,err:= os.OpenFile(saveFilePath,os.O_CREATE|os.O_RDWR,777)
	if err!=nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重试！")
		return
	}

	_ ,err = io.Copy(saveFile,file)
	if err!=nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重新尝试！")
		return
	}
	hash256 := sha256.New()
	filebytes ,_ := ioutil.ReadAll(file)
	hash256.Write(filebytes)
	hashbytes := hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashbytes))


	u.Ctx.WriteString("恭喜你,电子数据认证成功！！")
}