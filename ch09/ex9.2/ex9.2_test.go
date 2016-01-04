package main

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	ch := make(chan bool)

	go func() {
		fmt.Println(PopCount(0x1234567890ABCDEF))
		ch <- true
	}()

	go func() {
		fmt.Println(PopCount(0x1234567890ABCDEF))
		ch <- true
	}()

	<-ch
	<-ch

}
