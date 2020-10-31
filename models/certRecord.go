package models

import (
	"bytes"
	"encoding/gob"
)

/**
	该结构体用于定义链上数据保存的信息
 */
type CertRecord struct {
	CertId []byte//认证id，本身是md5值
	CertHash []byte//存证文件的hash值，本质是sha256值
	CertName string //认证人姓名
	Phone string//联系方式
	CertCard string//身份证号
	FileName string//文件的名称
	FileSize int64//文件的大小
	CertTime int64//认证的时间
}

/**
	序列化操作
 */
func (c CertRecord)Serialize()([]byte,error)  {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(),err
}

func DeserializeCertRecord(data []byte)(CertRecord,error)  {
	err :=
}