package function

import "fmt"

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
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

// 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。
func topoSort3(m map[string][]string) []string {
	visited := make(map[string]bool)
	var order []string
	seen := make(map[string]bool)
	var visitAll func(set map[string]bool, target string) error
	visitAll = func(set map[string]bool, target string) error {
		for item := range set {
			if visited[item] {
				return fmt.Errorf("%s need finish %s first however %[2]s need finish %[1]s first too", item, target)
			}
			visited[item] = true
			if !seen[item] {
				seen[item] = true
				s := slice2Set(m[item])
				if e := visitAll(s, item); e != nil {
					return e
				}

				order = append(order, item)
			}
		}
		return nil
	}
	for k := range m {
		e := visitAll(map[string]bool{k: true}, "")
		if e != nil {
			continue
		}
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
