package main

import (
	"time"
)

type Block struct {
	Timestamp     int64  // Timestamp of creation
	Data          []byte // Actual Data
	PrevBlockHash []byte // Hash of prev block
	Hash          []byte // This block's hash
	Nonce         int
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}
