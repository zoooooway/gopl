package multiplesort

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"text/tabwriter"
)

func TestMultiplesort(t *testing.T) {

	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
	}

	tb := Table{Tks: tracks}

	printTracks(tb.Tks)

	fmt.Println("--------------------------------------------")
	tb.AddSortKey("Title")
	tb.AddSortKey("Year")
	sort.Sort(tb)
	printTracks(tb.Tks)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
