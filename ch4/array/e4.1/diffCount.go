package main

import (
	"crypto/sha256"
	"fmt"
)

var pt [256]byte

func init() {
	for i, _ := range pt {
		pt[i] = pt[i/2] + byte(i&1)
	}
}

func main() {

	data1 := []byte("hash")
	data2 := []byte("dash")
	sha1 := sha256.Sum256(data1)
	sha2 := sha256.Sum256(data2)
	count := diffCount(sha1, sha2)
	fmt.Println(count)
}

func diffCount(sha1 [32]byte, sha2 [32]byte) uint8 {
	var c uint8
	for i := 0; i < 32; i++ {
		// 异或
		ret := sha1[i] ^ sha2[i]
		c += uint8(pt[ret])
	}
	return c
}
