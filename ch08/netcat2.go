// 客户端向server发送请求
// author: baoqiang
// time: 2019/3/20 下午7:18
package ch08

import (
	"net"
	"os"
	"log"
)

func RunClient2() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
