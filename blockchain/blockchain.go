package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
)

type Blockchain struct {
	Tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) CloseDb() {
	if err := bc.db.Close(); err != nil && err.Error() != "EOF" {
		log.Fatal(err)
	}
}

type ChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bci *ChainIterator) Next() *Block {
	var block *Block

	err := bci.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bci.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	bci.currentHash = block.PrevBlockHash

	if err != nil {
		log.Fatal(err)
	}

	return block
}

func (bc *Blockchain) Iterator() *ChainIterator {
	return &ChainIterator{bc.Tip, bc.db}
}

const dbFile = "db"
const blocksBucket = "blocks"

func NewBlockchain() *Blockchain {
	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesisBlock := NewBlock("Genesis Block", []byte{})
			log.Println("Creating new blockchain, because one does not exist...")

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Fatal(err)
			}

			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Fatal(err)
			}

			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Fatal(err)
			}

			tip = genesisBlock.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	if len(tip) == 0 {
		log.Fatal("no tip, after just creating blockchain")
	}

	bc := Blockchain{
		tip,
		db,
	}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b != nil {
			lastHash = b.Get([]byte("l"))
		}

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Fatal(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Fatal(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
