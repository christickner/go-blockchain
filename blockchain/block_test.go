package blockchain

import (
	"testing"
)

// NewBlock should generate a valid block with a Timestamp, Hash, and use the given Data and PrevBlockHash
func TestNewBlock(t *testing.T) {
	b := NewBlock("some data", []byte{'1', '2'})

	if string(b.Data) != "some data" {
		t.Error("data not in block")
	}

	if b.Timestamp == 0 {
		t.Error("timestamp was not set")
	}

	if b.PrevBlockHash[0] != '1' || b.PrevBlockHash[1] != '2' || len(b.PrevBlockHash) != 2 {
		t.Errorf("PrevBlockHash is not what we passed in, %v", b.PrevBlockHash)
	}

	if len(b.Hash) == 0 {
		t.Error("block hash was not set")
	}
}

func TestSerialization(t *testing.T) {
	b := NewBlock("new block", []byte{'1', '2'})
	s := b.Serialize()

	dBlock := DeserializeBlock(s)

	if string(dBlock.Hash) != string(b.Hash) {
		t.Errorf("wrong hash, got: %x, expected %x", dBlock.Hash, b.Hash)
	}

	if string(dBlock.Data) != "new block" {
		t.Errorf("wrong data, got: %s, expected %s", b.Data, dBlock.Data)
	}
}
