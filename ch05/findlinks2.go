// 函数多返回值
// author: baoqiang
// time: 2019/3/12 下午8:23
package ch05

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// ./main http://www.golang.org
func RunVisit2() {
	for _, url := range os.Args[1:] {
		links, err := findLinks2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "find links2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil

}
