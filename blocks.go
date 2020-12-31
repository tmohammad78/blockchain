package main

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp time.Time
	Data      []byte
	Hash      []byte
}

func (b *Block) string() string {
	return fmt.
		Sprintf("Timestamp : %s\n Data:%s\n",
			b.Timestamp, b.Data)
}
func NewBlock(data string) *Block {
	b := Block{
		Timestamp: time.Now(),
		Data:      []byte(data),
	}
	b.Hash = GenerateHash(b.Timestamp.UnixNano(), b.Data)
	return &b
}
