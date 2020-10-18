package utils

import (
	"io"
	"os"
)
//用来保存文件的函数
func SaveFile(fileName string,file io.Reader) (int64,error) {
	saveFile ,err:= os.OpenFile(fileName,os.O_CREATE|os.O_RDWR,777)
	if err!=nil {
		return -1,err
	}

	len,err := io.Copy(saveFile,file)
	if err!=nil {
		return -1,err
	}
	return len,nil
}
