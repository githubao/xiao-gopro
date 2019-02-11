// http服务
// author: baoqiang
// time: 2019/2/11 下午9:04
package ch01

import (
	"net/http"
	"github.com/labstack/gommon/log"
	"fmt"
)

func Server1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
