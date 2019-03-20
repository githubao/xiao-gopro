// 简化函数路由调用
// author: baoqiang
// time: 2019/3/18 下午3:43
package ch07

import (
	"log"
	"net/http"
)

func RunServer4() {
	db := database3{"shoes": 50, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
