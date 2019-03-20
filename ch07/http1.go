// 实现http服务
// author: baoqiang
// time: 2019/3/18 下午3:23
package ch07

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

func RunServer() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
