package models

type Userlogin struct {
	id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}
