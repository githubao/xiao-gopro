// defer关闭文件
// author: baoqiang
// time: 2019/3/13 上午11:21
package ch05

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// ./main https://golang.org/ https://golang.org/doc/gopher/frontpage.png
func RunFetch() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}
