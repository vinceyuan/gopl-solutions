package eval

import "fmt"

func (v Var) String() string {
	return "(" + string(v) + ")"
}

func (l literal) String() string {
	return fmt.Sprintf("(%f)", float64(l))
}

func (u unary) String() string {
	switch u.op {
	case '+':
		return "(+" + u.x.String() + ")"
	case '-':
		return "(-" + u.x.String() + ")"
	}
	return "???"
}

func (b binary) String() string {
	switch b.op {
	case '+':
		return "(" + b.x.String() + " + " + b.y.String() + ")"
	case '-':
		return "(" + b.x.String() + " - " + b.y.String() + ")"
	case '*':
		return "(" + b.x.String() + " * " + b.y.String() + ")"
	case '/':
		return "(" + b.x.String() + " / " + b.y.String() + ")"
	}
	return "(" + b.x.String() + "???" + b.y.String() + ")"
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return "pow(" + c.args[0].String() + ", " + c.args[1].String() + ")"
	case "sin":
		return "sin(" + c.args[0].String() + ")"
	case "sqrt":
		return "sqrt(" + c.args[0].String() + ")"
	}
	return "???()"
}
