package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.Add("Hello")
	fmt.Println(bc)
}
