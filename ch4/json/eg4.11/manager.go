package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const issueURL = "https://api.github.com/repos/%s/%s/issues"

const tempFilePattern = "issueTempFile"

var tempFileDir, _ = os.UserHomeDir()

type Response struct {
	Id     uint
	Title  string
	Body   string
	URL    string
	Number uint16
	State  string
	Labels []*Label
}

type Label struct {
	Id          uint
	NodeId      string `json:"node_id"`
	URL         string
	Name        string
	Description string
}

type Request struct {
	Id     uint16 `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	URL    string `json:"url"`
	Number uint16 `json:"number"`
	State  string `json:"state"`
}

var token string

var cli = &http.Client{}

// 编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。
func main() {
	fmt.Fprintln(os.Stdout, "please enter your authorization token before you use this application: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "illegal input")
		os.Exit(1)
	}
	token = scanner.Text()

	fmt.Fprintln(os.Stdout, "now you can use this application: ")
	fmt.Fprintln(os.Stdout, "use '--m create -o <owner> -r <repo>' to create an issue")
	fmt.Fprintln(os.Stdout, "use '--m get -o <owner> -r <repo> -i <issue_number>' to get an issue")
	fmt.Fprintln(os.Stdout, "use '--m update -o <owner> -r <repo> -i <issue_number>' to update an issue")
	fmt.Fprintln(os.Stdout, "use '--m close -o <owner> -r <repo> -i <issue_number>' to close an issue")
	fmt.Fprintln(os.Stdout, "enter ':exit' to exit")

	for {
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "illegal input1")
			continue
		}

		cmd := scanner.Text()
		pts := strings.Split(cmd, " ")
		var method, owner, repo, issueNumber string
		for i := 0; i < len(pts); i++ {
			if strings.HasPrefix(pts[i], "--m") {
				method = pts[i+1]
				i++
			} else if pts[i] == "-o" {
				owner = pts[i+1]
				i++
			} else if pts[i] == "-r" {
				repo = pts[i+1]
				i++
			} else if pts[i] == "-i" {
				issueNumber = pts[i+1]
				i++
			} else if pts[i] == ":exit" {
				os.Exit(1)

			} else {
				fmt.Fprintln(os.Stderr, "illegal input2")
				break
			}
		}

		r, e := managerIssues(method, owner, repo, issueNumber)
		if e != nil {
			fmt.Fprintln(os.Stderr, e.Error())
			continue
		}

		body := r.Body

		if r.StatusCode < 200 || r.StatusCode >= 300 {
			fmt.Fprintln(os.Stderr, r.Status)
			body.Close()
			continue
		}

		if r.StatusCode == 204 {
			fmt.Fprintln(os.Stdin, r.Status)
			body.Close()
			continue
		}

		var result Response

		if e := json.NewDecoder(body).Decode(&result); e != nil {
			fmt.Fprintln(os.Stderr, e.Error())
			continue
		}

		data, e := json.MarshalIndent(result, "", "    ")
		if e != nil {
			log.Fatalf("JSON marshaling failed: %s", e)
		}

		fmt.Printf("%s\n", data)
		body.Close()

	}

}

func managerIssues(method string, owner string, repo string, issueNumber string) (resp *http.Response, err error) {
	var data string
	if method == "get" {
		url := fmt.Sprintf(issueURL, owner, repo) + "/" + issueNumber
		return getIssue(url)
	} else if method == "create" {
		url := fmt.Sprintf(issueURL, owner, repo)
		// via default text editor
		if data, e := openTextEditorToWriter(tempFileDir, tempFilePattern); e != nil {
			return resp, e
		} else {
			return createIssue(url, data)
		}

	} else if method == "update" {
		url := fmt.Sprintf(issueURL, owner, repo) + "/" + issueNumber
		// via default text editor
		if data, e := openTextEditorToWriter(tempFileDir, tempFilePattern); e != nil {
			return resp, e
		} else {
			return updateIssue(url, data)
		}

	} else if method == "close" {
		url := fmt.Sprintf(issueURL, owner, repo) + "/" + issueNumber + "/lock"
		return closeIssue(url, data)
	} else {
		return nil, fmt.Errorf("illegal input")
	}

}

func getIssue(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	rep, e := cli.Do(req)
	return rep, e
}

func createIssue(url string, data string) (resp *http.Response, err error) {
	reader := bytes.NewReader([]byte(data))
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	rep, e := cli.Do(req)
	return rep, e
}

func updateIssue(url string, data string) (resp *http.Response, err error) {
	reader := bytes.NewReader([]byte(data))
	req, e := http.NewRequest(http.MethodPatch, url, reader)
	if e != nil {
		return nil, e

	}
	req.Header.Add("Authorization", "Bearer "+token)

	rep, e := cli.Do(req)
	return rep, e
}

func closeIssue(url string, data string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	rep, e := cli.Do(req)
	return rep, e
}

func openTextEditorToWriter(dir, pattern string) (data string, e error) {
	if fs, e := os.CreateTemp(dir, pattern); e != nil {
		return data, e
	} else {
		var req Request
		json, e := json.MarshalIndent(&req, "", "    ")
		if e != nil {
			log.Fatalf("JSON marshaling failed: %s", e)
		}

		fs.Write([]byte(json))

		name := fs.Name()
		fs.Close()

		cmd := exec.Command("notepad", name)
		if e := cmd.Run(); e != nil {
			os.Remove(fs.Name())
			return data, e
		}

		if bytes, e := os.ReadFile(name); e != nil {
			return data, e
		} else {
			if e := os.Remove(name); e != nil {
				fmt.Fprintln(os.Stderr, e.Error())
			}
			return string(bytes), nil
		}
	}
}
