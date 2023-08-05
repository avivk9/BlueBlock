package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] // last block
	newBlock := newBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock) // append new block
}

func NewGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}} // Creates a new Blockchain, with a list of Block pointers and a first genesis block
}
