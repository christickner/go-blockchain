package blockchain

import "testing"

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()

	if genesisBlock := bc.blocks[0]; genesisBlock == nil {
		t.Error("The new Blockchain must have the genesis Block")
	}

	if string(bc.blocks[0].Data) != "Genesis Block" {
		t.Error("Unexpected data in the genesis Block")
	}
}
