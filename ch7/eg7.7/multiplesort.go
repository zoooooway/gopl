package multiplesort

import (
	"time"

	"golang.org/x/exp/slices"
)

// 很多图形界面提供了一个有状态的多重排序表格插件：
// 主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
// 定义一个sort.Interface的实现用在这样的表格中。
// 比较这个实现方式和重复使用sort.Stable来排序的方式。

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type Table struct {
	Tks     []*Track
	orderby []string
}

func (x *Table) AddSortKey(k string) {
	// ignore duplicates key
	if slices.Contains(x.orderby, k) {
		return
	}
	x.orderby = append(x.orderby, k)
}

func (x *Table) ClearSortKey() {
	x.orderby = []string{}
}

func (x Table) Len() int { return len(x.Tks) }
func (x Table) Less(i, j int) bool {
	for _, k := range x.orderby {
		switch k {
		case "Title":
			if x.Tks[i].Title == x.Tks[j].Title {
				continue
			}
			return x.Tks[i].Title < x.Tks[j].Title
		case "Artist":
			if x.Tks[i].Artist == x.Tks[j].Artist {
				continue
			}
			return x.Tks[i].Artist < x.Tks[j].Artist
		case "Album":
			if x.Tks[i].Album == x.Tks[j].Album {
				continue
			}
			return x.Tks[i].Album < x.Tks[j].Album
		case "Year":
			if x.Tks[i].Year == x.Tks[j].Year {
				continue
			}
			return x.Tks[i].Year < x.Tks[j].Year
		case "Length":
			if x.Tks[i].Length == x.Tks[j].Length {
				continue
			}
			return x.Tks[i].Length < x.Tks[j].Length

		}
	}

	return x.Tks[i].Title < x.Tks[j].Title
}

func (x Table) Swap(i, j int) { x.Tks[i], x.Tks[j] = x.Tks[j], x.Tks[i] }
