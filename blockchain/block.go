package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
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
}
//创建一个新区块
func NewBlock(height int64,prevHash []byte,data []byte) Block {
	block := Block{
		Height:    height+1,
		TimeStemp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		//Hash:      nil,
		Version:   "0x01",
	}
	var sumBytes []byte
	heightByte,_ := utils.Int64ToByte(block.Height)
	timeStempByte,_ := utils.Int64ToByte(block.TimeStemp)
	versionByte := utils.StringToBytes(block.Version)
	sumBytes = bytes.Join([][]byte{
		heightByte,timeStempByte,block.PrevHash,block.Data,versionByte},[]byte{})
	//block.Hash
	block.Hash = utils.SHA256HashBlock(sumBytes)
	return block
}

func CreateGenesisBlock() Block {
	blockGenesis :=NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return blockGenesis
}