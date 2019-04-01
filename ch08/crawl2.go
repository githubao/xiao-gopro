// 并发限制数量
// author: baoqiang
// time: 2019-04-01 21:35
package ch08

import (
	"fmt"
	"github.com/githubao/xiao-gopro/ch05"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	list, err := ch05.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func RunCrawl2() {
	worklist := make(chan []string)

	var n int
	n++

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}

}
