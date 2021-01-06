package main

import (
	"crypto/sha256"
	"fmt"
)

func EasyHash(data ...interface{}) []byte {
	hasher := sha256.New()
	fmt.Fprint(hasher, data...)
	return hasher.Sum(nil)
}
