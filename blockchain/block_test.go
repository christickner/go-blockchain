package blockchain

import (
	"testing"
)

// NewBlock should generate a valid block with a Timestamp, Hash, and use the given Data nd PrevBlockHash
func TestNewBlock(t *testing.T) {
	b := NewBlock("some data", []byte{'1', '2'})

	if string(b.Data) != "some data" {
		t.Fail()
	}

	if b.Timestamp == 0 {
		t.Fail()
	}

	if b.PrevBlockHash[0] != '1' || b.PrevBlockHash[1] != '2' || len(b.PrevBlockHash) != 2 {
		t.Fail()
	}

	if len(b.Hash) == 0 {
		t.Fail()
	}
}
