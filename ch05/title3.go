// recover特定的类型错误
// author: baoqiang
// time: 2019/3/13 上午11:36
package ch05

import (
	"fmt"

	"net/http"
	"strings"

	"os"

	"golang.org/x/net/html"
)

// ./main https://golang.org/ https://golang.org/doc/gopher/frontpage.png

// title: multiple title elements
// title: https://golang.org/doc/gopher/frontpage.png has type image/png, not text/html
func RunTitle3() {
	for _, arg := range os.Args[1:] {
		if err := title3(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			//pass
		case bailout{}:
			//expected
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	visitFunc := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}

	forEachNode(doc, visitFunc, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

func title3(url string) error {
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

	title, err := soleTitle(doc)
	if err != nil {
		return err
	}

	fmt.Println(title)
	return nil
}
