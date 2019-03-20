// 客户端连接
// author: baoqiang
// time: 2019/3/20 下午7:05
package ch08

import (
	"io"
	"log"
	"net"
	"os"
)

func RunClient1() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
