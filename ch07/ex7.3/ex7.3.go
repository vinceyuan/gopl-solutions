package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
func (t *tree) String() string {
	var str string

	var visit func(tr *tree)

	visit = func(tr *tree) {
		if tr.left != nil {
			visit(tr.left)
		}
		str = fmt.Sprintf("%s %d", str, tr.value)
		if tr.right != nil {
			visit(tr.right)
		}
	}
	visit(t)
	return str
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root.String())
}
