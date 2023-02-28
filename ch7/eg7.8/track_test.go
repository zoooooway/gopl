package track

import (
	"fmt"
	msort "gopl/ch7/eg7.7"
	"net/http"
	"sort"
	"testing"
)

var tracks = []*msort.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, msort.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, msort.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, msort.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, msort.Length("4m24s")},
}

var tb msort.Table = msort.Table{Tks: tracks}

func TestTrack(t *testing.T) {

	http.HandleFunc("/", handler)
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	key := r.Form.Get("key")
	if len(key) > 0 {
		tb.AddSortKey(key)
		sort.Sort(tb)
	}
	HtmlTracks(w, tb.Tks)
}
