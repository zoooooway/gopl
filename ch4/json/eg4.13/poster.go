package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

const POSTER_API = "http://www.omdbapi.com/?apikey=[appkey]&t=%s"

type Response struct {
	Title  string
	Poster string
}

func main() {
	fmt.Println("Please enter movie title to search poster:")
	s := bufio.NewScanner(os.Stdin)
	for {
		if s.Scan() {
			title := strings.ToLower(s.Text())
			res, e := getPoster(title)
			if e != nil {
				fmt.Println(e.Error())
				continue
			}

			fmt.Printf("title: %s, \nposter:%s\n", res.Title, res.Poster)
		}
	}

}

func getPoster(p string) (Response, error) {
	url := fmt.Sprintf(POSTER_API, html.EscapeString(p))
	r, e := http.Get(url)
	if e != nil {
		return Response{}, e
	}

	if r.StatusCode != http.StatusOK {
		r.Body.Close()
		return Response{}, errors.New(r.Status)
	}

	var res Response
	e = json.NewDecoder(r.Body).Decode(&res)
	r.Body.Close()
	if e != nil {
		return Response{}, e
	}

	return res, nil

}
