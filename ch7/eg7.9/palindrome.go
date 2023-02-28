package palindrome

import (
	"sort"
)

// sort.Interface类型也可以适用在其它地方。
// 编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。
// 假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。

func IsPalindrome(s sort.Interface) bool {
	for i, l := 0, s.Len()-1; i <= l-i; i++ {
		if !(!s.Less(i, l-i) && !s.Less(l-i, i)) {
			return false
		}
	}
	return true
}
