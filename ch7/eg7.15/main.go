package main

import (
	"bufio"
	"fmt"
	eval "gopl/ch7/eg7.13"
	"log"
	"os"
	"strconv"
)

// eg7.15： 编写一个从标准输入中读取一个单一表达式的程序，用户及时地提供对于任意变量的值，然后在结果环境变量中计算表达式的值。
// 优雅的处理所有遇到的错误。
func main() {
	os.Stdout.Write([]byte("enter an expression：\n"))
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	expr, e := eval.Parse(input)
	if e != nil {
		log.Fatal(e.Error())
	}
	vs := make(map[eval.Var]bool)
	expr.Check(vs)
	env := eval.Env{}
	for k, _ := range vs {
		fmt.Fprintf(os.Stdout, "enter %s = ", k)
		s.Scan()
		v, e := strconv.ParseFloat(s.Text(), 64)
		if e != nil {
			log.Fatal("sorry, the input is illegal", e.Error())
		}
		env[k] = v
	}
	fmt.Fprintf(os.Stdout, "value = %f", expr.Eval(env))
}
