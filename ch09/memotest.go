// 工具函数
// author: baoqiang
// time: 2019-04-08 21:12
package ch09

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingUrls() <-chan string {
	ch := make(chan string)

	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}

		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

func Sequential(m M) {
	for url := range incomingUrls() {
		start := time.Now()

		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(m M) {
	var n sync.WaitGroup

	for url := range incomingUrls() {
		n.Add(1)

		go func(url string) {
			defer n.Done()

			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}

			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}
