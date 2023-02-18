package parser

import (
	html "gopl/ch5/func/eg5.7"
	"log"
	"testing"
)

func TestParser(t *testing.T) {
	h := `<!-- saved from url=(0037)https://dictionary.cambridge.org/zhs/ -->	<html>		<head id="header">		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">		<meta name="color-scheme" content="light dark">	</head>		<body>		<!-- hello world -->		<div id="haha1" class="line-gutter-backdrop"></div>		<img src="//www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png" />		<form autocomplete="off"><label id="haha2" class="line-wrap-control">自动换行<input type="checkbox"					aria-label="自动换行"></label>		</form>		<form autocomplete="off"><label id="haha1" class="line-wrap-control">自动换行<input type="checkbox"			aria-label="自动换行"></label>		<div id="ggg" class="line-gutter-backdrop"></div>	</form>		</body>		</html>`
	r := NewReader(h)
	doc, err := r.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}
	html.Print(doc)
}
