package blockchain

import (
	"DataCertPlatform/utils"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Target *big.Int//系统的目标值
	Block Block//要找到nonce值对应的区块
}

func NewPow(block Block)  {
	t := big.NewInt(1)
	t = t.Lsh(t,255)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	var i int64
	for  i=0; ;i++  {
		if  {
			
		}
		bytestring := utils.SHA256Pow(pow.Block.Hash,i)
		fmt.Println("区块和nonce的hash",bytestring)
	}
}
