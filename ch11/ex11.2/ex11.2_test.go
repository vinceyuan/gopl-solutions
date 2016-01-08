package intset

import "testing"

func TestExample1(t *testing.T) {
	var result string
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	result = x.String()
	if result != "{1 9 144}" {
		t.Errorf("x Add 1, 144, 9 = %q", result)
	}

	y.Add(9)
	y.Add(42)
	result = y.String()
	if result != "{9 42}" {
		t.Errorf("y Add 9, 42 = %q", result)
	}

	x.UnionWith(&y)
	result = x.String()
	if result != "{1 9 42 144}" {
		t.Errorf("x.UnionWith(&y) = %q", result)
	}

	if x.Has(9) != true {
		t.Errorf("x.Has(9) = %q", x.Has(9))
	}

	if x.Has(123) != false {
		t.Errorf("x.Has(123) = %q", x.Has(123))
	}

}
