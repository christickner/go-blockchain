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
	if err := bc.db.Close(); err != nil {
		log.Print("close ")
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
		log.Print("view ")
		log.Fatal(err)
	}

	return nil
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
		log.Print("could not open ")
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesisBlock := NewBlock("Genesis Block", []byte{})

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Print("could not create bucket: ")
				log.Fatal(err)
			}

			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Printf("%x coult not PUT genesis hash\n ", genesisBlock.Hash)
				log.Fatal(err)
			}

			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Printf("%x coult not PUT genesis L\n ", genesisBlock.Hash)
				log.Fatal(err)
			}

			tip = genesisBlock.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		if err != nil {
			log.Print("unknown err: ")
			log.Fatal(err)
		}

		return nil
	})

	log.Printf("%x created tip\n", tip)

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
			log.Printf("%x = L, FOUND last hash when creating new block", lastHash)
		}

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Print("coult not put hash ")
			log.Fatal(err)
		}

		log.Printf("%x put block\n", newBlock.Hash)

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Print("x ")
			log.Fatal(err)
		}

		log.Printf("%x put l as\n", newBlock.Hash)

		bc.Tip = newBlock.Hash

		log.Printf("%x bc\n", bc.Tip)

		return nil
	})

	if err != nil {
		log.Print("n ")
		log.Fatal(err)
	}
}
