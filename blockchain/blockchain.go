package blockchain

import (
	"github.com/bolt"
)

/**
区块链结构体的定义，代表的是整条区块链
 */
const BUCKET_NAME  = "blocks"
const LAST_HASH  = "lasthash"
const BLOCKCHAIN_NAME  = "blockchain.db"

type BlockChain struct {
	LashHash []byte//表示区块链中最新区块的哈希，用于查找最新的区块内容
	BoltDb *bolt.DB//区块链中操作区块数据文件的数据库操作对象
}
/**
创建一条区块链
 */
func NewBlockChain() BlockChain {
	//创世区块
	genesis :=CreateGenesisBlock()
	db, err := bolt.Open(BLOCKCHAIN_NAME,0600,nil)
	if err != nil {
		panic(err.Error())
	}
	bc := BlockChain{
		LashHash: genesis.Hash,
		BoltDb:   db,
	}
	//将创世区块放入文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket ,err :=tx.CreateBucket([]byte(BUCKET_NAME))
		if err!= nil {
			panic(err.Error())
		}
		//序列化
		genesisbytes := genesis.Serialize()
		//把创世区块放入桶中
		bucket.Put(genesis.Hash,genesisbytes)
		//更新最新区块的hash值，方便查找
		bucket.Put([]byte(LAST_HASH),genesis.Hash)
		return nil
	})
	return bc
}
//保存数据到区块链中: 先生成一个新区块,然后将新区块添加到区块链中
func (bc BlockChain)SaveBlock(date []byte)  {
	//1、从文件当中读取到最新的区块
	db := bc.BoltDb
	var lastblock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockbytes := bucket.Get(lastHash)
		//反序列化
		lastblock,_= DeSerialize(lastBlockbytes)
		return nil
	})
	newBlock := NewBlock(lastblock.Height+1,lastblock.Hash,date)
	//把新区块存放到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//序列化
		newBlockbytes := newBlock.Serialize()
		//存储新区块
		bucket.Put(newBlock.Hash,newBlockbytes)
		//更新lasthash的值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		return nil
	})
}
