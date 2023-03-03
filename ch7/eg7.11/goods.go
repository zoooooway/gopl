package goods

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Dollars float32

func (d Dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Database map[string]Dollars

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// eg7.12 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。html/template包（§4.6）可能会对你有帮助

var templ = `<h1>goods</h1>
<table style='width: 20%'>
<tr style='text-align: left'>
  <th>item</th>
  <th>price</th>
</tr>
{{range $key, $value := . }}
<tr>
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`
var table = template.Must(template.New("goods").Parse(templ))

func (db Database) Table(w http.ResponseWriter, req *http.Request) {
	table.Execute(w, db)
}

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db Database) Update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	p := req.URL.Query().Get("price")
	if len(p) == 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is not valid: %q\n", item)
		return
	}

	pf, e := strconv.ParseFloat(p, 32)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is not valid: %q\n", item)
		return
	}
	price := Dollars(pf)
	db[item] = price
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db Database) Create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item: %q already exists\n", item)
		return
	}

	p := req.URL.Query().Get("price")
	if len(p) == 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price cannot empty: %q\n", item)
		return
	}

	pf, e := strconv.ParseFloat(p, 32)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is not valid: %f\n", pf)
		return
	}
	price := Dollars(pf)
	db[item] = price
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db Database) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s\n", "success")
}
