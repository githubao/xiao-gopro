// 函数作为参数传递，遍历网页节点
// author: baoqiang
// time: 2019/3/12 下午8:43
package ch05

import (
	"net/http"
	"os"

	"fmt"

	"golang.org/x/net/html"
)

func RunOutline2() {
	for _, url := range os.Args[1:] {
		outline2(url)
	}
}

func outline2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

// .* The precision is not specified in the format string,
// but as an additional integer value argument preceding the argument that has to be formatted.
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
