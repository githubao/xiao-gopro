// 注册对应的方法
// author: baoqiang
// time: 2019/3/18 下午3:38
package ch07

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer3() {
	db := database3{"shoes": 50, "socks": 5}

	mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price))

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type database3 map[string]dollars

func (db database3) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database3) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
