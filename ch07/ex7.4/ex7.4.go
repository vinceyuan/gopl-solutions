package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type StringReader string

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = len(*r)
	copy(p, []byte(*r))
	err = io.EOF // Must set EOF, otherwise it does not end
	return
}

func NewReader(str string) *StringReader {
	var s StringReader
	s = StringReader(str)
	return &s
}

func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
