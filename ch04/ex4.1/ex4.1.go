package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [8]byte

func init() {
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
	}
}

func main() {
	hash1 := sha256.Sum256([]byte("hello world"))
	hash2 := sha256.Sum256([]byte("hello worlD"))
	printHash(hash1)
	printHash(hash2)
	fmt.Println(diffBits(hash1, hash2))
}

func printHash(hash [32]byte) {
	for _, v := range hash {
		fmt.Printf("%X", v)
	}
	fmt.Println()
}

func diffBits(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		byte1 := hash1[i]
		byte2 := hash2[i]
		for j := 0; j < 8; j++ {
			if byte1&pc[j] == byte2&pc[j] {
				count++
			}
		}
	}
	return count
}
