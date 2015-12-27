package main

import "fmt"

func noReturn() {
	panic(5)
}

func main() {
	defer func() {
		val := recover()
		fmt.Println(val)
	}()
	noReturn()
}
