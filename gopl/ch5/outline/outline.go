package outline

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 第一层循环时传入的始终是最初的 stack
		// stack只在递归内部修改
		Outline(stack, c)
	}
}

// CountNodeNames implements exercise 5.2s
func CountNodeNames(counter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CountNodeNames(counter, c)
	}
}

func ShowText(n *html.Node) {
	if n == nil {
		return
	}

	// 处理文本节点：只输出非空内容
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Println(text)
		}

	}

	// 只有非 script/style 元素才递归其子节点
	if n.Type == html.ElementNode {
		lowerData := strings.ToLower(n.Data)
		if lowerData != "script" && lowerData != "style" {
			ShowText(n.FirstChild)
		}
	} else {
		ShowText(n.FirstChild)
	}

	// 总是递归下一个兄弟（无论当前节点类型）
	ShowText(n.NextSibling)
}
