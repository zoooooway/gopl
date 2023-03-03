package eval

import (
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return " " + string(u.op) + " " + u.x.String()
}

func (b binary) String() string {
	return "(" + b.x.String() + " " + string(b.op) + " " + b.y.String() + ")"
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return c.args[0].String() + "^" + c.args[1].String()
	case "sin":
		return "sin(" + c.args[0].String() + ")"
	case "sqrt":
		return c.args[0].String() + "^" + "(1/2)"
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (m min) String() string {
	return "min(" + m.args[0].String() + ", " + m.args[1].String() + ")"
}
