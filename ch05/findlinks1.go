// 递归visit节点
// author: baoqiang
// time: 2019/3/5 上午11:54
package ch05

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

// /
// /
// #
// /doc/
// /pkg/
// /project/
// /help/
// #
// /dl/
// https://developers.google.com/site-policies#restrictions
// /LICENSE
// /doc/tos.html
// http://www.google.com/intl/en/policies/privacy/

func RunVisit() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find links1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links

}
