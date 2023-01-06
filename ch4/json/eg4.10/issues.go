package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

const Placeholder = "-------------------------------------------------------------------------------------"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
func main() {

	m := make(map[string][]*Issue)

	if res, err := searchIssuesWithDate(os.Args[1:], ">=", time.Now().AddDate(0, 0, -30)); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	} else {
		m["last month"] = res.Items
	}

	if res, err := searchIssuesWithDate(os.Args[1:], ">=", time.Now().AddDate(0, -6, 0)); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	} else {
		m["last half year"] = res.Items
	}

	if res, err := searchIssuesWithDate(os.Args[1:], ">=", time.Now().AddDate(-1, 0, 0)); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	} else {
		m["last year"] = res.Items
	}

	if res, err := searchIssuesWithDate(os.Args[1:], "<", time.Now().AddDate(-1, 0, 0)); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	} else {
		m["more then one year"] = res.Items
	}

	for k, v := range m {
		fmt.Printf("%s: %s\n", k, Placeholder)
		for _, item := range v {
			fmt.Printf("#%-5d\t%-20.20s\t%.55s\t%v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}

func searchIssuesWithDate(query []string, flag string, date time.Time) (*IssuesSearchResult, error) {
	query = append(query, "created:"+flag+date.Format("2006-01-02"))
	return searchIssues(query)
}

func searchIssues(query []string) (*IssuesSearchResult, error) {
	q := strings.Join(query, " ")
	res, err := http.Get(IssuesURL + "?q=" + url.QueryEscape(q))
	if err != nil {
		fmt.Printf("error while searching for issues: %s", err.Error())
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	body := res.Body
	var issues IssuesSearchResult
	if err := json.NewDecoder(body).Decode(&issues); err != nil {
		fmt.Printf("decode failed: %s", err.Error())
		return nil, err
	}
	res.Body.Close()
	return &issues, nil
}
