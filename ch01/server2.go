// http并发请求锁限制
// author: baoqiang
// time: 2019/2/11 下午9:07
package ch01

import (
	"net/http"
	"code.byted.org/gopkg/pkg/log"
	"sync"
	"fmt"
)

var mu sync.Mutex
var count int

func Server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/favicon.ico", handlerFavicon)
	//http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println(count)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handlerFavicon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello.ico")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d\n", count)
	mu.Unlock()
}
