package models

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/utils"
	"fmt"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

/**
* 将用户信息保存到数据库中去的函数
 */
func (u User)AddUser() (int64, error) {
	//1、将密码进行hash计算，得到密码hash值，然后在存
	u.Password = utils.MD5HashString(u.Password)
	//execute， .exe
	result, err :=db_mysql.Db.Exec("insert into user_message(user_phone,user_pwd)" +
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
func (u User)QueryUser()(*User,error) {
	u.Password = utils.MD5HashString(u.Password)
	row :=db_mysql.Db.QueryRow("select user_phone from user_message where user_phone = ? and user_pwd = ?",
		u.Phone,u.Password)
	err := row.Scan(&u.Phone)
	if err!=nil {
		return nil,err
	}
	return &u,nil 
}

//根据phone查询用户id
func (u User)QueryUserIdByPhone() (*User,error) {
	row := db_mysql.Db.QueryRow("select id from user_message where user_phone = ?",u.Phone)
	err := row.Scan(&u.Id)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return &u,nil
}
