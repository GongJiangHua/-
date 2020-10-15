package models

import "github.com/GongJiangHua/HelloBeego_01/db_mysql"

/**
上传文件的记录
 */
type UploadRecord struct {
	Id         int
	UserId     int
	FileName   string
	FileSize   int
	FileCert   string
	FileTitle  string
	CertTime   int

}

func (u UploadRecord) SaveRecord()  {
	db_mysql.DB.Exec("insert into uploadrecord " +
		"(user_id,file_name,file_size,file_sert,file_title,cert_time)" +
		"values (?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)
}
