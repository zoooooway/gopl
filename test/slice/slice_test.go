package slice

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	// words := []string{"ss", "dd"}
	// var nw []string
	// nw = append(nw, words...)
	// fmt.Println(nw)
	// fmt.Println(words)

	// words[1] = "halo"
	// fmt.Println(nw)
	// fmt.Println(words)

	// ss := make([]string, 10)
	// ns := make([]string, 10)
	// ss[0] = "en"
	// ss[1] = "zh"
	// ns = append(ns, ss...)
	// fmt.Println(ss)
	// fmt.Println(ns)

	// ss[1] = "halo"
	// fmt.Println(ss)
	// fmt.Println(ns)
	// Ensome{}.write()
	var a A
	var w io.Writer = &bytes.Buffer{}
	fmt.Printf("a=(%T, %v)\n", a, a)
	fmt.Printf("w=(%T, %v)\n", w, w)

	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	sort.Sort(byArtist(tracks))
	i := sort.Reverse(byArtist(tracks))

	fmt.Print(i)
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type A struct {
	int
	io.Writer
}

// func (a A) write() {
// 	fmt.Println("test write")
// }
