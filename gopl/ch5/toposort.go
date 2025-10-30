package ch5

import "sort"

// 给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；
// 我们的目标是选择一组课程，这组课程必须确保按书序血虚时，能全部被完成。
// 这类问题被称作拓扑排序，从概念上说，前置条件可以构成有向图
// 图中的顶点表示课程，边表示课程间的依赖关系。显然，图中应该无环，也即是说，
// 从某点出发的边，最终不会回到该点。
var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":   {"discrete math"},
	"databases":         {"data structures"},
	"discrete math":     {"intro to programming"},
	"formal languages":  {"discrete math"},
	"networks":          {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {
		"data structures",
		"computer organization",
	},
}

// 使用深度优先搜索整张图，获得了符合要求的课程序列。
func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// 当匿名函数需要被递归调用时，必须先声明一个变量，再将匿名函数赋值给这个变量。
	// 如果不分成两步，函数字面量无法与 visitAll 绑定，就无法递归调用该匿名函数。
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

	// 对  prereqs 的 key 值进行排序，保证每次运行 toposort 都以相同的遍历顺序遍历
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
