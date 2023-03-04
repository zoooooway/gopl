package main

import (
	"fmt"
	eval "gopl/ch7/eg7.13"
	"html/template"
	"net/http"
	"strconv"
)

var templ = `
<!DOCTYPE html>
<html>
<body>

<h1>simple calculator</h1>
expression: <input type="text" id="expr" placeholder="input expression">
<br>
<button type="button" onclick="eval()">calculate</button>
<br>
result: <span id="result">hello<span>

<script>
function eval()
{
	var xmlhttp;
	if (window.XMLHttpRequest)
	{
		//  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
		xmlhttp=new XMLHttpRequest();
	}
	else
	{
		// IE6, IE5 浏览器执行代码
		xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
	}
	xmlhttp.onreadystatechange=function()
	{
		if (xmlhttp.readyState==4 && xmlhttp.status==200)
		{
			document.getElementById("result").innerText=xmlhttp.responseText;
		}
	}
	xmlhttp.open("GET", "http://localhost/eval?expr=" + encodeURIComponent(document.getElementById('expr').value), true);
	xmlhttp.send();
}
</script>

</body>
</html>

`

var calc = template.Must(template.New("calculator").Parse(templ))

var env = eval.Env{}

// eg7.16： 编写一个基于web的计算器程序。
func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(templ)) })
	http.HandleFunc("/eval", handleEval)
	http.ListenAndServe("localhost:80", nil)
}

func handleEval(w http.ResponseWriter, r *http.Request) {
	expr := r.URL.Query().Get("expr")
	if len(expr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "illegal expression: %s", expr)
		return
	}

	ep, e := eval.Parse(expr)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "illegal expression: %s", e.Error())
		return
	}

	if e := ep.Check(make(map[eval.Var]bool)); e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "illegal expression: %s", e.Error())
		return
	}

	result := strconv.FormatFloat(float64(ep.Eval(env)), 'f', -1, 64)
	fmt.Fprint(w, result)
}
