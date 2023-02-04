package function

import (
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。
func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(set map[string]bool)
	visitAll = func(set map[string]bool) {
		for item := range set {
			if !seen[item] {
				seen[item] = true
				s := slice2Set(m[item])
				visitAll(s)
				order = append(order, item)
			}
		}
	}
	for k := range m {
		visitAll(map[string]bool{k: true})
		break
	}

	return order
}

func slice2Set(item []string) map[string]bool {
	set := make(map[string]bool)
	for _, v := range item {
		set[v] = true
	}
	return set
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func isVaildTopo(lesson []string, m map[string][]string) bool {
	seen := map[string]bool{}

	var isAllSeen func(lesson []string) bool

	isAllSeen = func(lesson []string) bool {
		for _, v := range lesson {
			pre := m[v]
			if len(pre) == 0 {
				seen[v] = true
				continue
			}

			for _, l := range pre {
				if !seen[l] {
					return false
				}
			}
			seen[v] = true
		}
		return true
	}

	return isAllSeen(lesson)
}
