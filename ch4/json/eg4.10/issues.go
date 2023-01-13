package main

import (
	"encoding/json"
	"fmt"
	"gopl/ch4/json/eg4.10/github"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
func main() {

	m := make(map[string][]*github.Issue)

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
		fmt.Printf("%s: %s\n", k, github.Placeholder)
		for _, item := range v {
			fmt.Printf("#%-5d\t%-20.20s\t%.55s\t%v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}

func searchIssuesWithDate(query []string, flag string, date time.Time) (*github.IssuesSearchResult, error) {
	query = append(query, "created:"+flag+date.Format("2006-01-02"))
	return searchIssues(query)
}

func searchIssues(query []string) (*github.IssuesSearchResult, error) {
	q := strings.Join(query, " ")
	res, err := http.Get(github.IssuesURL + "?q=" + url.QueryEscape(q))
	if err != nil {
		fmt.Printf("error while searching for issues: %s", err.Error())
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	body := res.Body
	var issues github.IssuesSearchResult
	if err := json.NewDecoder(body).Decode(&issues); err != nil {
		fmt.Printf("decode failed: %s", err.Error())
		return nil, err
	}
	res.Body.Close()
	return &issues, nil
}
