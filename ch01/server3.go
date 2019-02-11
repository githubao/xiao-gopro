// http请求头
// author: baoqiang
// time: 2019/2/11 下午9:53
package ch01

import (
	"net/http"
	"github.com/labstack/gommon/log"
	"fmt"
)

//curl
//GET / HTTP/1.1
//Header["User-Agent"] = ["curl/7.54.0"]
//Header["Accept"] = ["*/*"]
//Host = "localhost:8000"
//RemoteAddr = "127.0.0.1:56855"
//
//chrome
//GET / HTTP/1.1
//Header["Upgrade-Insecure-Requests"] = ["1"]
//Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36"]
//Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"]
//Header["Accept-Encoding"] = ["gzip, deflate, br"]
//Header["Accept-Language"] = ["zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7"]
//Header["Connection"] = ["keep-alive"]
//Host = "localhost:8000"
//RemoteAddr = "127.0.0.1:56859"

func Server3() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
