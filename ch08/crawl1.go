// 并发无限制的爬虫
// author: baoqiang
// time: 2019-04-01 21:30
package ch08

import (
	"fmt"
	"github.com/githubao/xiao-gopro/ch05"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)

	list, err := ch05.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func RunCrawl1() {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}
