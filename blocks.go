package main

import (
	"bytes"
	"fmt"
	"time"
)

type Block struct {
	Timestamp time.Time
	Data      []byte
	prevHash  []byte
	Hash      []byte
}

func (b *Block) String() string {
	return fmt.
		Sprintf("Timestamp : %s\n Data:%s\n Hash:%x\n prevHash:%x\n----\n",
			b.Timestamp, b.Data, b.Hash, b.prevHash)
}

func (b *Block) Validate() error {
	h := EasyHash(b.Timestamp.UnixNano(), b.Data, b.prevHash)
	if !bytes.Equal(h, b.Hash) {
		return fmt.Errorf("the hash is invalid it should be %x is %x", h, b.Hash)
	}
	return nil
}
func NewBlock(data string, prevHash []byte) *Block {
	b := Block{
		Timestamp: time.Now(),
		Data:      []byte(data),
		prevHash:  prevHash,
	}
	b.Hash = EasyHash(b.Timestamp.UnixNano(), b.Data, b.prevHash)
	return &b
}

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) Add(data string) {
	ln := len(bc.Blocks)
	if ln == 0 {
		panic("panic")
	}
	bc.Blocks = append(
		bc.Blocks, NewBlock(data, bc.Blocks[ln-1].Hash))
}
func (b *BlockChain) String() string {
	var ret string
	for i := range b.Blocks {
		ret += b.Blocks[i].String()
	}
	return ret
}

func (bc *BlockChain) Validate() error {
	for i := range bc.Blocks {
		if err := bc.Blocks[i].Validate(); err != nil {
			return fmt.Errorf("Block chain is not valid : %w", err)
		}
		if i == 0 {
			continue
		}
		if !bytes.Equal(bc.Blocks[i].prevHash, bc.Blocks[i-1].Hash) {
			return fmt.Errorf("the order is invalid it  %x should be  %x", bc.Blocks[i].prevHash, bc.Blocks[i-1].Hash)
		}
	}
	return nil
}
func NewBlockChain() *BlockChain {
	bc := BlockChain{}
	bc.Blocks = []*Block{NewBlock("Genesis Block", []byte{})}
	return &bc
}
