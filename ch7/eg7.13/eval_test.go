package eval

import (
	"fmt"
	"log"
	"testing"
)

func TestEval(t *testing.T) {
	input := "sqrt(x*30)+68/22-min(2,19)+sin(pow(5,2))*23"
	expr, e := Parse(input)
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Println(expr.String())
	fmt.Println(expr.Eval(Env{}))
}
