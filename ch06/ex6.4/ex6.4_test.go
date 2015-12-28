package intset

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func TestExample2(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestExample3(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	if x.String() != "{1 9 144}" {
		t.Errorf("x.String() should be {1 9 144}")
	}
	fmt.Println(x.Len())
	if x.Len() != 3 {
		t.Errorf("x.Len() should be 3")
	}
	x.Remove(9)
	fmt.Println(x.String())
	if x.String() != "{1 144}" {
		t.Errorf("x.String() should be {1 144}")
	}
	x.Clear()
	fmt.Println(x.Len())

	x.Add(1)
	x.Add(144)
	x.Add(9)

	y := x.Copy()
	fmt.Println(y.String())
	x.Remove(9)
	fmt.Println(x.String())
	fmt.Println(y.String())

	x.AddAll(2, 3, 4)
	fmt.Println("x", x.String())

	y.Add(166)
	fmt.Println("y", y.String())
	x.IntersectWith(y)
	fmt.Println(x.String())
	if x.String() != "{1 144}" {
		t.Errorf("x.String() should be {1 144}")
	}

	var a, b IntSet
	a.AddAll(12, 20)
	b.AddAll(12, 30)
	a.DifferenceWith(&b)
	fmt.Println("a", a.String())

	var c, d IntSet
	c.AddAll(12, 20)
	d.AddAll(12, 30)
	c.SymmetricDifferenceWith(&d)
	fmt.Println("c", c.String())

	for i, val := range c.Elems() {
		fmt.Println(i, val)
	}
}
