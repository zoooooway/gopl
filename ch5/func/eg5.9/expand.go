package expand

import "strings"

// 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
