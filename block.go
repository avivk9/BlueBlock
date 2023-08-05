package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  // Timestamp of creation
	Data          []byte // Actual Data
	PrevBlockHash []byte // Hash of prev block
	Hash          []byte // This block's hash
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))                       // Convert timestamp to bytes
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // Combine headers
	hash := sha256.Sum256(headers)                                                // Hash headers together

	b.Hash = hash[:]
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
