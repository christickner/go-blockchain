package blockchain

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestNewProofOfWork(t *testing.T) {
	b := NewBlock("some data", []byte{})

	pow := NewProofOfWork(b)

	if pow.block != b {
		t.Error("block should be the same")
	}
}

func TestRunProofOfWork(t *testing.T) {
	b := NewBlock("Chris sent 1 coin", []byte{})
	pow := NewProofOfWork(b)

	// mine the block
	goldenNonce, minedHash := pow.Run()

	// verify the pow found the goldenNonce that produce a valid minedHash of the block data
	blockData := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			intToHex(pow.block.Timestamp),
			intToHex(int64(targetBits)),
			intToHex(int64(goldenNonce)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(blockData)
	expectedHash := hash[:]

	if string(expectedHash) != string(minedHash) {
		t.Errorf("pow hash \"%x\" does not match expected \"%x\"", minedHash, expectedHash)
	}
}
