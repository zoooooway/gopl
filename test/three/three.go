package main

import (
	"errors"
	"fmt"
)

var str = "hello"
var e = errors.New("new error")

func main() {
	e := t("world")
	fmt.Printf("%s", e)
}

func t(str string) (e error) {
	fmt.Println(str)
	return e
}
