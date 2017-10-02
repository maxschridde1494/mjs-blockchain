package chain

import (
	"blockchain/block"
	"log"

	"github.com/boltdb/bolt"
)

type Blockchain struct {
	Tip []byte
	Db  *bolt.DB
}

const blockBucket = "BlockBucket"

// var BC *Blockchain

// func Setup() {
// 	if BC != nil {
// 		log.Println("Blockchain is already setup")
// 		return
// 	}
// 	BC = NewBlockchain()
// 	log.Println("Blockchain is initialized")
// }

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Println("Error in AddBlock grabbing last hash.")
	} else {
		newBlock := block.NewBlock(data, lastHash)
		err = bc.Db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockBucket))
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Println("Error in AddBlock serialization")
				return err
			}
			err = b.Put([]byte("l"), newBlock.Hash)
			bc.Tip = newBlock.Hash
			log.Println("Added new block with data: %s", data)
			return nil
		})
		if err != nil {
			log.Println("Error in AddBlock db.Update")
		}
	}
}

func NewGenesisBlock() *block.Block {
	log.Println("Creating Genesis Block")
	return block.NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, _ := bolt.Open("mjs-blockchain.db", 0600, nil)

	_ = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, _ := tx.CreateBucket([]byte(blockBucket))
			_ = b.Put(genesis.Hash, genesis.Serialize())
			_ = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	bc := Blockchain{tip, db}
	return &bc
}

//============BLOCK CHAIN ITERATOR============
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.Db}
	return bci
}

func (i *BlockchainIterator) Next() *block.Block {
	var blk *block.Block
	_ = i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		blk = block.DeserializeBlock(encodedBlock)

		return nil
	})
	i.currentHash = blk.PrevBlockHash
	return blk
}
