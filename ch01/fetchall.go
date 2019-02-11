// 并发http请求
// author: baoqiang
// time: 2019/2/11 下午7:05
package ch01

import (
	"time"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"os"
)

// 0.66s    5749 https://golang.org
// 0.70s    6805 https://godoc.org
// 1.17s    4154 http://gopl.io
// 1.17s elapsed
func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
