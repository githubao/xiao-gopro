// 广度优先遍历
// author: baoqiang
// time: 2019/3/12 下午9:15
package ch05

import (
	"fmt"
	"log"
	"os"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist

		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}

	}

}

func crawl(url string) []string {
	fmt.Println(url)

	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func RunCrawl() {
	breadthFirst(crawl, os.Args[1:])
}
