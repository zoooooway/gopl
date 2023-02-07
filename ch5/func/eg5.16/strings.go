package strings

import "strings"

// 编写多参数版本的strings.Join
func Join(seq string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	var b strings.Builder
	b.WriteString(elems[0])
	for _, v := range elems[1:] {
		b.WriteString(seq)
		b.WriteString(v)
	}
	return b.String()
}
