// 打印节点信息
// author: baoqiang
// time: 2019/3/5 下午12:27
package ch05

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

func RunOutline() {
	filename := "/Users/baoqiang/go_repos/src/github.com/githubao/xiao-gopro/ch05/golangorg.txt"
	f,_ := os.Open(filename)

	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find links1: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}
