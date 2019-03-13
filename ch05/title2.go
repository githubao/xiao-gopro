// 使用defer
// author: baoqiang
// time: 2019/3/13 上午11:12
package ch05

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// ./main https://golang.org https://golang.org/doc/gopher/frontpage.png

// The Go Programming Language
// submit search
// title: https://golang.org/doc/gopher/frontpage.png has type image/png, not text/html
func RunTitle2() {
	for _, arg := range os.Args[1:] {
		if err := title2(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func title2(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)

	return nil
}
