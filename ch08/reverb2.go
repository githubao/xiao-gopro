// 并发处理客户端过来的请求
// author: baoqiang
// time: 2019/3/20 下午7:33
package ch08

import (
	"net"
	"time"
	"log"
	"bufio"
)

func RunServer4() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn4(conn)
	}

}
func handleConn4(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}