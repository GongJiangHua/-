package blockchain

import (
	"bytes"
	"encoding/gob"
	_ "github.com/astaxie/beego/cache"
	"time"
)

/**
定义区块结构体，用于表示区块
 */
type Block struct {
	Height int64//表示区块的高度，第几个区块
	TimeStemp int64//区块产生的时间戳
	PrevHash []byte//前一个字段的hash
	Data []byte//数据字段
	Hash []byte//当前字段的hash值
	Version string//版本号
	Nonce int64//区块对应的nonce值
}
//创建一个新区块
func NewBlock(height int64,prevHash []byte,data []byte) Block {
	block := Block{
		Height:    height,
		TimeStemp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		//Hash:      ,
		Version:   "0x01",
	}
	pow := NewPow(block)
	hashBlock,nonce := pow.Run()
	block.Hash=hashBlock
	block.Nonce=nonce
	//将block结构体数据转化为[]byte类型
	//var sumBytes []byte
	//heightByte,_ := utils.Int64ToByte(block.Height)
	//timeStempByte,_ := utils.Int64ToByte(block.TimeStemp)
	//versionByte := utils.StringToBytes(block.Version)
	//nonceByte,_ := utils.Int64ToByte(nonce)
	////bytes.jion拼接
	//sumBytes = bytes.Join([][]byte{
	//	heightByte,timeStempByte,block.PrevHash,block.Data,versionByte,nonceByte},[]byte{})
	//block.Hash
	//调用hash计算，对区块进行sha256哈希值计算
	//block.Hash = utils.SHA256HashBlock(sumBytes)
	//挖矿竞争
	return block
}

func CreateGenesisBlock() Block {
	blockGenesis :=NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return blockGenesis
}

/**
对区块进行序列化
 */
func (b Block)Serialize() []byte {
	buff := new(bytes.Buffer)//缓冲区
	encoder := gob.NewEncoder(buff)
	encoder.Encode(b)//将区块b放入到缓冲区中
	return buff.Bytes()
}
//反序列化将硬盘中的区块内容放入内存中
func DeSerialize(d []byte) (*Block,error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil,err
	}
	return &block,nil
}