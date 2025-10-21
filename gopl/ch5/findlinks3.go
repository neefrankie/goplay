package ch5

import (
	"fmt"
	"log"
)

func BreadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// f(item)... 将f返回的一组元素一个个添加到worklist中
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
