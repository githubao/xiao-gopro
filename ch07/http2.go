// 实现路由
// author: baoqiang
// time: 2019/3/18 下午3:28
package ch07

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:8080/price?item=socks
func RunServer2() {
	db := database2{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type database2 map[string]dollars

func (db database2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}

}
