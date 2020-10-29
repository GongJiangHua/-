package blockchain

import (
	"github.com/bolt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"math/big"
)

/**
区块链结构体的定义，代表的是整条区块链
 */
const BUCKET_NAME  = "blocks"
const LAST_HASH  = "lasthash"
const BLOCKCHAIN_NAME  = "blockchain.db"

var CHAIN *BlockChain

type BlockChain struct {
	LashHash []byte//表示区块链中最新区块的哈希，用于查找最新的区块内容
	BoltDb *bolt.DB//区块链中操作区块数据文件的数据库操作对象
}
/**
创建一条区块链
 */
func NewBlockChain() BlockChain {
	var bc BlockChain
	//先打开文件
	db, err := bolt.Open(BLOCKCHAIN_NAME,0600,nil)
	//查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("BUCKET_NAME"))
		if bucket==nil {
			bucket ,err = tx.CreateBucket([]byte("BUCKET_NAME"))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_HASH))
		if len(lastHash)==0 {//没有lasthash,弄一个创世区块
			//创世区块
			genesis :=CreateGenesisBlock()
			genesisBytes := genesis.Serialize()
			bucket.Put(genesis.Hash,genesisBytes)
			bucket.Put([]byte(LAST_HASH),genesis.Hash)
			bc  = BlockChain{
				LashHash: genesis.Hash,
				BoltDb:   db,
			}
		}else {
			lasthash1 := bucket.Get([]byte(LAST_HASH))
			bc = BlockChain{
				LashHash: lasthash1,
				BoltDb:   db,
			}
		}
		return nil
	})
	return bc
}


//保存数据到区块链中: 先生成一个新区块,然后将新区块添加到区块链中
func (bc *BlockChain)SaveBlock(date []byte) (Block,error) {
	//1、从文件当中读取到最新的区块
	db := bc.BoltDb
	var lastblock *Block
	//error的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			err = errors.New("桶中的数据为空")
			return err
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockbytes := bucket.Get(bc.LashHash)
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
		//把新创建的区块存储到文件中
		bucket.Put(newBlock.Hash,newBlockbytes)
		//更新lasthash的值，更新为最新存储的区块的区块hash值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		bc.LashHash = newBlock.Hash
		return nil
	})
	//返回值语句，newBlock，err，其中err可能包含错误信息
	return newBlock,err
}

func (bc BlockChain)QueryAllBlocks()([]*Block,error)  {
	blocks := make([]*Block,0)//blocks是一个切片容器，用于存放查询到的区块
	db := bc.BoltDb
	var err error
	//从chaindb文件查询所有的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			err = errors.New("查询区块链失败！")
			return err
		}
		eachHash := bc.LashHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		//bucket存在
		for {
			lastBlockBytes := bucket.Get([]byte(eachHash))
			lastBlock,_:=DeSerialize(lastBlockBytes)
			blocks=append(blocks, lastBlock)
			eachBig.SetBytes(lastBlock.Hash)
			if zeroBig.Cmp(eachBig)==0 {
				break
			}
			eachHash = lastBlock.PrevHash
		}
		return nil
	})
	return blocks,err
}

/**
该方法用于完成根据用户输入的区块高度查询对应的区块信息
 */
func (bc BlockChain)QueryBlockByHeight(height int64) (*Block,error) {
		//获取db文件
	db:= bc.BoltDb
	//定义一个err错误
	if height<0 {
		return nil,nil
	}
	var errs error
	var eachBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			errs = errors.New("读取桶中数据失败！！")
			return errs
		}
		eachHash := bc.LashHash
		for{
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock,errs = DeSerialize(eachBlockBytes)
			if errs!= nil {
				return errs
			}
			if eachBlock.Height<height {
				break
			}
			if height==eachBlock.Height {
				break
			}
			eachHash= eachBlock.PrevHash
		}
		return nil
	})
	return eachBlock,errs
}
