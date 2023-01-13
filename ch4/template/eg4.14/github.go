package main

import (
	"bufio"
	"fmt"
	"gopl/ch4/json/eg4.10/github"
	"net/http"
	"os"
)

const ISSUES_URL = "https://api.github.com/issues"

var cli = &http.Client{}

var token string

type Response struct {
	Title string
	Body  string
	User  github.User
}

// 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
func main() {
	fmt.Fprintln(os.Stdout, "please enter your authorization token before you use this application: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "illegal input")
		os.Exit(1)
	}
	token = scanner.Text()

	listIssues()
}

func listIssues() (resp *http.Response, err error) {
	req, e := http.NewRequest(http.MethodGet, ISSUES_URL, nil)
	if e != nil {
		return nil, e

	}
	req.Header.Add("Authorization", "Bearer "+token)

	rep, e := cli.Do(req)
	return rep, e
}
