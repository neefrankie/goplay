package ch5

import (
	"fmt"
	"sort"
)

func TopoSort1(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(item string)
	visitAll = func(item string) {
		for _, dep := range m[item] {
			if !seen[dep] {
				seen[dep] = true
				visitAll(dep)
				order = append(order, dep)
			}
		}
	}

	for key := range m {
		if !seen[key] {
			seen[key] = true
			visitAll(key)
			order = append(order, key)
		}
	}

	return order
}

var PrereqsCycle = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
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

func TopoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	onStack := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				onStack[item] = true

				visitAll(m[item])
				onStack[item] = false
				order = append(order, item)
			} else if onStack[item] {
				fmt.Printf("cycle detected: %s\n", item)
				continue
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
