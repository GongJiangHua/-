package db_mysql

import (
	"DataCertPlatform/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
/**
 * 连接mysql数据库
 */
func Connect(){

	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassword)

	//连接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {// err不为nil，表示连接数据库时出现了错误, 程序就在此中断就可以，不用再执行了。
		//早解决，早解决
		panic("数据库连接错误，请检查配置")
	}
	Db = db
	fmt.Println(db)
}
/**
* 将用户信息保存到数据库中去的函数
 */
func AddUser(u models.User)(int64, error){
	//1、将密码进行hash计算，得到密码hash值，然后在存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	psswordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(psswordBytes)
	//execute， .exe
	result, err :=Db.Exec("insert into user_message(user_phone,user_pwd)" +
		" values(?,?) ", u.Phone,u.Password)
	if err != nil {
		return -1,err
	}
	row,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return row,nil
}

//登录页面用户登录进行数据查询
func QueryUser(u models.Userlogin)(int,error ) {
	//1、将密码进行hash计算，得到密码hash值，然后进行比较
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	psswordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(psswordBytes)
	row :=Db.QueryRow("select password from user_message where use_phone = ?",
		u.Phone)
	var password int
	err := row.Scan(&password)
	if err!=nil {
		return 0,err
	}
	return password,nil
}
