package method

import (
	"bytes"
	"fmt"
	"math/bits"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/BIT, uint(x%BIT)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/BIT, uint(x%BIT)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BIT; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", BIT*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// eg6.1 +++
// 为bit数组实现下面这些方法

// return the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, v := range s.words {
		count += bits.OnesCount(v)
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	// position x
	i, v := x/BIT, x%BIT
	s.words[i] ^= (0x8000000000000000 >> v)
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	is := IntSet{}
	is.words = append(is.words, s.words...)

	return &is
}

// eg6.1 ---

// eg6.2+++
// 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
func (s *IntSet) AddAll(vs ...int) {
	for _, v := range vs {
		s.Add(v)
	}
}

// eg6.2---

// eg6.3+++
// (*IntSet).UnionWith会用|操作符计算两个集合的并集，
// 我们再为IntSet实现另外的几个函数:
// IntersectWith（交集：元素在A集合B集合均出现），
// DifferenceWith（差集：元素出现在A集合，未出现在B集合），
// SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	var re IntSet
	re.words = make([]uint, max(len(s.words), len(s.words)))

	for i, v := range t.words {
		re.words[i] = s.words[i] & v
	}
	return &re
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var re IntSet
	re.words = make([]uint, max(len(s.words), len(s.words)))

	for i, v := range t.words {
		re.words[i] = s.words[i] & (-v - 1)
	}
	return &re
}

func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var re IntSet
	re.words = make([]uint, max(len(s.words), len(s.words)))

	for i, v := range t.words {
		re.words[i] = s.words[i] ^ v
	}
	return &re
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// eg6.3---

// eg6.4+++
// 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作
func (s IntSet) Elems() (els []int) {
	for i, v := range s.words {
		if v == 0 {
			continue
		}
		for j := 0; j < BIT; j++ {
			if (1 << j & v) != 0 {
				els = append(els, i*BIT+j)
			}
		}
	}
	return
}

// eg6.4---

// eg6.5+++
// 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。
// 当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，
// 这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)

const (
	BIT = 32 << (^uint(0) >> 63)
)

// eg6.5---
