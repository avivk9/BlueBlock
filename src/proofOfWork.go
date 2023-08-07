package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 16 // How hard will it be to mine

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // Upper bounder of the proofs, proof < targe => proof is valid.

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte { // combining the increasement and the data
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions { // iterate all transactoins
		txHashes = append(txHashes, tx.ID) // appends the ID of each transaction
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{})) // sum all transactions tpgether

	return txHash[:] // convert the array to slices
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64

	fmt.Printf("Mining a new block")
	for nonce < maxNonce {
		data := pow.prepareData(nonce) // data prepared with nonce
		hash = sha256.Sum256(data)     // get the hash
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:]) // convert hash to big.Int

		if hashInt.Cmp(pow.target) == -1 { // check if hash is valid pow (proof of work)
			break // stop when valid
		} else {
			nonce++ // increase when not valid
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
