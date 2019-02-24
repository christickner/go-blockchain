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

func TestAddBlock(t *testing.T) {
	bc := NewBlockchain()

	bc.AddBlock("the second block")

	if len(bc.blocks) != 2 {
		t.Errorf("Expected 2 Blocks in the Blockchain, got: %d", len(bc.blocks))
	}

	if string(bc.blocks[1].Data) != "the second block" {
		t.Errorf("Unexpected Block data, got: %s", string(bc.blocks[1].Data))
	}
}

func TestBlocks(t *testing.T) {
	bc := NewBlockchain()

	bc.AddBlock("block 2")

	if len(bc.blocks) != 2 {
		t.Errorf("Expected 2 Blocks, got: %d", bc.Blocks())
	}

	for i, b := range bc.Blocks() {
		if b != bc.blocks[i] {
			t.Errorf("Block at index %d is not expected", i)
		}
	}
}
