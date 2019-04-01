// 限制协程的数量
// author: baoqiang
// time: 2019-04-01 21:41
package ch08

import "os"

func RunCrawl3() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i <= 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}
