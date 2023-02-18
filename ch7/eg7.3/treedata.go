package tree

import (
	"bytes"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// 为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法
// 展示tree类型的值序列。
// 使用先序遍历
func (t *tree) String() string {
	var buf bytes.Buffer
	buf.WriteRune('{')
	if t != nil {
		buf.WriteString(strconv.FormatInt(int64(t.value), 10))
		recursion(t.left, &buf)
		recursion(t.right, &buf)
	}

	buf.WriteRune('}')
	return buf.String()
}

func recursion(t *tree, buf *bytes.Buffer) {
	if t != nil {
		buf.WriteRune(rune(' '))
		buf.WriteString(strconv.FormatInt(int64(t.value), 10))
		recursion(t.left, buf)
		recursion(t.right, buf)
	}
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
