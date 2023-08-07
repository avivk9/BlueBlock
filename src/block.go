package main

import (
	"time"
)

type Block struct {
	Timestamp     int64          // Timestamp of creation
	Transactions  []*Transaction // Every block stores transuctions
	PrevBlockHash []byte         // Hash of prev block
	Hash          []byte         // This block's hash
	Nonce         int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
