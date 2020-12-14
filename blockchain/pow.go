package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"fmt"
	"math/big"
)

const DIFFICULTY  = 5
type ProofOfWork struct {
	Target *big.Int//系统的目标值
	Block Block//要找到nonce值对应的区块
}

func NewPow(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t,255-DIFFICULTY)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}

/**
用于寻找合适的nonce值
 */
func (p ProofOfWork)Run() ([]byte,int64) {
	var nonce int64
	nonce = 0
	var blockHash []byte
	for{
		block := p.Block
		heightBytes,_ := utils.Int64ToByte(block.Height)
		timeStempBytes,_ := utils.Int64ToByte(block.TimeStemp)
		versionByte := utils.StringToBytes(block.Version)
		nonceBytes,_ := utils.Int64ToByte(nonce)
		//已有区块信息和尝试的nonce值的拼接信息
		blockBytes := bytes.Join([][]byte{
			heightBytes,timeStempBytes,block.PrevHash,block.Data,versionByte,nonceBytes,
		},[]byte{})
		blockHash = utils.SHA256HashBlock(blockBytes)
		target := p.Target//目标值
		var hashBig *big.Int//声明和定义
		hashBig = new(big.Int)//分配内存空间，为变量分配地址
		hashBig = hashBig.SetBytes(blockHash)
		fmt.Println("当前的nonce值:",nonce)
		if hashBig.Cmp(target)==-1 {
			//结束执行，退出
			break
		}
		nonce++
	}
	//将找到的符合规则的nonce值返回
	return blockHash,nonce
}
//