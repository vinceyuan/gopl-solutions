package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface {
	String() string
} // CharData or *Element

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	childrenStr := ""
	for _, child := range e.Children {
		childrenStr += child.String()
	}
	return "<" + e.Type.Local + ">" + childrenStr + "</" + e.Type.Local + ">"
}

func main() {
	//fmt.Println("Input xml and press ctrl-D to end")
	//dec := xml.NewDecoder(os.Stdin)
	input := "<A><B><C>hello</C><D>abc</D></B><C>world</C></A>"
	fmt.Println("Parsing input:", input)
	dec := xml.NewDecoder(strings.NewReader(input))
	var root Node
	var stack [][]Node // stack of nodes
	depth := 0
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			root = stack[0][len(stack[0])-1]
			stacklen := len(stack)
			if depth < stacklen-1 {
				if element, ok := stack[stacklen-2][len(stack[stacklen-2])-1].(Element); ok {
					element.Children = stack[stacklen-1]
				}
				stack = stack[:stacklen-1] // pop
			}
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "7.18: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			depth++
			if len(stack) < depth {
				nodes := make([]Node, 0)
				stack = append(stack, nodes)
			}
			node := &Element{tok.Name, tok.Attr, []Node{}}
			stack[depth-1] = append(stack[depth-1], node)
		case xml.EndElement:
			depth--
			stacklen := len(stack)
			if depth < stacklen-1 {
				if element, ok := stack[stacklen-2][len(stack[stacklen-2])-1].(*Element); ok {
					element.Children = stack[stacklen-1]
				}
				stack = stack[:stacklen-1] // pop
			}
		case xml.CharData:
			stacklen := len(stack)
			if element, ok := stack[stacklen-1][len(stack[stacklen-1])-1].(*Element); ok {
				element.Children = []Node{CharData(string(tok))}
			}
		}
	}

	fmt.Println("\nOutput:", root.String())
}
