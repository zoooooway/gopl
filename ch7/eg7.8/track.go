package track

import (
	"fmt"
	msort "gopl/ch7/eg7.7"
	"html/template"
	"io"
	"os"
	"text/tabwriter"
)

// 使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。
// 将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

var templ = `<h1>tracks</h1><a href='http://localhost:8080/tracks/sort/clear'>清除排序</a>
<table>
<tr style='text-align: left'>
  <th><a href='http://localhost:8080/tracks/sort?key=Title'>Title</a></th>
  <th><a href='http://localhost:8080/tracks/sort?key=Artist'>Artist</a></th>
  <th><a href='http://localhost:8080/tracks/sort?key=Album'>Album</a></th>
  <th><a href='http://localhost:8080/tracks/sort?key=Year'>Year</a></th>
  <th><a href='http://localhost:8080/tracks/sort?key=Length'>Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`

var table = template.Must(template.New("trackstable").Parse(templ))

func HtmlTracks(w io.Writer, tracks []*msort.Track) {
	table.Execute(w, tracks)
}

func printTracks(tracks []*msort.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
