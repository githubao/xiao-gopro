// http请求
// author: baoqiang
// time: 2019/2/11 下午7:00
package ch01

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

// https://golang.org http://gopl.io https://godoc.org
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
}
