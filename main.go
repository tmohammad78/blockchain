package main

import (
	"fmt"
	"log"
)

func main() {
	bc := NewBlockChain()
	bc.Add("Hello")
	bc.Add("Welcome")
	if err := bc.Validate(); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(bc)
}
