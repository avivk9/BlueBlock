package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

const dbFile = "database/BlueBlocksDB.db" // the name of the BoltDB file
const blocksBucket = "blocks"             // The name of the bucket to store blockchain blocks in BoltDB

func IntToHex(n int64) []byte {
	hexString := fmt.Sprintf("%x", n)
	return []byte(hexString)
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
