package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopl/ch4/json/eg4.10/github"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const ISSUES_URL = "https://api.github.com/issues"

var cli = &http.Client{}

var token string

var templ = `{{. | len}} issues:
{{range .}}----------------------------------------
Number: {{.Number}}
Title: {{.Title}}
Body: {{.Body}}
State: {{.State}}
CreateAt: {{.CreatedAt}}
{{if .User}}User: 
    Login: {{.User.Login}}{{end}}
{{if .Milestone}}Milestone: 
    Title: {{.Milestone.Title}}
    Description: {{.Milestone.Description}}
    State: {{.Milestone.State}}
    OpenIssues: {{.Milestone.OpenIssues}}
    ClosedIssues: {{.Milestone.ClosedIssues}}
    CreatedAt: {{.Milestone.CreatedAt | daysAgo}} days{{end}}
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

// 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
func main() {
	fmt.Fprintln(os.Stdout, "please enter your authorization token before you use this application: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("illegal input")
	}
	token = scanner.Text()

	r, e := listIssues()
	if e != nil {
		r.Body.Close()
		log.Fatal(e.Error())
	}

	if r.StatusCode != http.StatusOK {
		r.Body.Close()
		log.Fatal(e.Error())
	}
	var issues []github.Issue
	e = json.NewDecoder(r.Body).Decode(&issues)
	fmt.Printf("%v", issues)
	r.Body.Close()
	if e != nil {
		log.Fatal(e.Error())
	}

	if e := report.Execute(os.Stdout, &issues); e != nil {
		log.Fatal(e.Error())
	}

}

func listIssues() (*http.Response, error) {
	req, e := http.NewRequest(http.MethodGet, ISSUES_URL, nil)
	if e != nil {
		return nil, e

	}
	req.Header.Add("Authorization", "Bearer "+token)

	rep, e := cli.Do(req)
	return rep, e
}
