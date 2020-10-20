package utils

import (
	"bytes"
	"encoding/binary"
)

/**
将一个int64转化为[]byte
 */
func Int64ToByte(num int64) ([]byte,error) {
	//Buffer:缓冲区，增益
	buff := new(bytes.Buffer)
	//buff.Write()通过一系列的Write方法向缓冲区写入数据
	//buff.Bytes()通过Bytes方法从缓冲区中获取数据
	/**
	 *
	 */
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		return nil,err
	}
	//从缓冲区读取数据
	return buff.Bytes(),err
}

func StringToBytes(str string)[]byte  {
	return []byte(str)
}
