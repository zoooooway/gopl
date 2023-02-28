package main

import (
	"fmt"
	msort "gopl/ch7/eg7.7"
	track "gopl/ch7/eg7.8"
	"log"
	"net/http"
	"sort"
)

var tracks = []*msort.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, msort.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, msort.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, msort.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, msort.Length("4m24s")},
}

var tb msort.Table = msort.Table{Tks: tracks}

func main() {
	http.HandleFunc("/tracks/sort", handler)
	http.HandleFunc("/tracks/sort/clear", handleClear)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	fmt.Println(key)
	if len(key) > 0 {
		tb.AddSortKey(key)
		sort.Sort(tb)
	}
	writeTable(w)
}

// handler echoes the Path component of the request URL r.
func handleClear(w http.ResponseWriter, _ *http.Request) {
	tb.ClearSortKey()
	writeTable(w)
}

func writeTable(w http.ResponseWriter) {
	track.HtmlTracks(w, tb.Tks)
}
