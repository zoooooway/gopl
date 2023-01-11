package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const OFFLINE_FILE = "./comics.txt"

var comics []Comic

const (
	CURRENT_COMIC_URL = "https://xkcd.com/info.0.json"
	GET_COMIC_URL     = "https://xkcd.com/%d/info.0.json"
)

type Comic struct {
	Month      string
	Num        uint
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

// 流行的web漫画服务xkcd也提供了JSON接口。
// 例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。
// 下载每个链接（只下载一次）然后创建一个离线索引。
// 编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
func main() {
	fmt.Println("Please enter keyword to search comic:")
	s := bufio.NewScanner(os.Stdin)
	for {
		if s.Scan() {
			keyword := strings.ToLower(s.Text())
			cs := search(keyword)
			if len(*cs) == 0 {
				fmt.Println("Sorry, we couldn't find anything. Try another keyword.")
				continue
			}
			fmt.Println("We found the following:")
			for _, c := range *cs {
				fmt.Printf("title:\n %s,\nurl:\n %s \n\n", c.Title, c.Img)
			}
		}
	}

}

// init offline data
// only update new data
func init() {
	f, e := os.Stat(OFFLINE_FILE)
	if e != nil {
		if _, e := os.Create(OFFLINE_FILE); e != nil {
			log.Fatalln(e.Error())
		}
		f, e = os.Stat(OFFLINE_FILE)
		if e != nil {
			log.Fatalln(e.Error())
		}
	}

	data, e := os.ReadFile(f.Name())
	if e != nil {
		log.Fatalln(e.Error())
	}

	if len(data) > 0 {
		e = json.Unmarshal(data, &comics)
		if e != nil {
			log.Fatalln(e.Error())
		}
	}

	var start uint = 1
	if len(comics) > 0 {
		sort.Slice(comics, func(i, j int) bool {
			return comics[i].Num < comics[j].Num
		})

		start = comics[len(comics)-1].Num
	}
	fmt.Printf("The latest number of offline data is: %d\n", start)

	var curr Comic
	// get latest num
	resp, e := http.Get(CURRENT_COMIC_URL)
	if e != nil {
		log.Fatalln(e.Error())
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalln(e.Error())
	}
	if e := json.NewDecoder(resp.Body).Decode(&curr); e != nil {
		resp.Body.Close()
		log.Fatalln(e.Error())
	}
	end := curr.Num
	fmt.Printf("The latest number data is: %d\n", end)

	if start < end {
		// update new data
		ch := make(chan *Comic)
		d := end - start
		for i := uint(1); i <= d; i++ {
			go queryComic(start+i, ch)
		}
		for i := uint(1); i <= d; i++ {
			r := <-ch
			if r != nil {
				comics = append(comics, *r)
			}
		}
	}

	bytes, e := json.Marshal(&comics)
	if e != nil {
		log.Fatalln(e.Error())
	}
	os.WriteFile(f.Name(), bytes, os.ModeAppend)
}

func search(keyword string) *[]Comic {
	var r []Comic
	for _, c := range comics {
		lt := strings.ToLower(c.Title)
		if strings.Contains(lt, keyword) {
			r = append(r, c)
		}
	}
	return &r
}

func queryComic(num uint, ch chan<- *Comic) {
	fmt.Printf("try to get data of num: %d \n", num)
	c := http.Client{}
	c.Timeout = 20 * time.Second
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(GET_COMIC_URL, num), nil)
	c.Do(r)
	resp, e := c.Do(r)
	if e != nil {
		fmt.Printf("%s getting data, number is: %d \n", e.Error(), num)
		ch <- nil
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("%s getting data, number is: %d \n", resp.Status, num)
		ch <- nil
		return
	}

	var item Comic
	if e := json.NewDecoder(resp.Body).Decode(&item); e != nil {
		resp.Body.Close()
		fmt.Printf("%s getting data, number is: %d \n", e.Error(), num)
		ch <- nil
		return
	}
	fmt.Printf("get data of num: %d finish \n", num)
	ch <- &item
}
